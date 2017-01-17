#include "thumbnailer.h"
#include <magick/pixel_cache.h>
#include <string.h>

const unsigned long thumbWidth = 150, thumbHeight = 150;

int thumbnail(const void *src,
			  const size_t size,
			  struct Thumbnail *thumb,
			  ExceptionInfo *ex)
{
	ImageInfo *info = NULL;
	Image *img = NULL, *sampled = NULL, *scaled = NULL;
	double scale;
	int err = 0;

	// Read image
	info = CloneImageInfo(NULL);
	GetExceptionInfo(ex);
	// Read only the first frame/page of GIFs and PDFs
	info->subimage = 0;
	info->subrange = 1;
	img = BlobToImage(info, src, size, ex);
	if (img == NULL) {
		goto end;
	}

	// Image already fits thumbnail
	if (img->columns <= thumbWidth && img->rows <= thumbHeight) {
		err = writeThumb(img, thumb, ex);
		goto end;
	}

	// Maintain aspect ratio
	if (img->columns >= img->rows) {
		scale = (double)(img->columns) / (double)(thumbWidth);
	} else {
		scale = (double)(img->rows) / (double)(thumbHeight);
	}
	const int width = (unsigned long)(img->columns / scale);
	const int height = (unsigned long)(img->rows / scale);

	// Subsample to 4 times the thumbnail size. A decent enough compromise
	// between quality and performance for images arround the thumbnail size
	// and much bigger ones.
	sampled = SampleImage(img, width * 4, height * 4, ex);
	if (sampled == NULL) {
		goto end;
	}

	// Scale to thumbnail size
	scaled = ResizeImage(sampled, width, height, BoxFilter, 1, ex);
	if (scaled == NULL) {
		goto end;
	}

	err = writeThumb(scaled, thumb, ex);

end:
	if (img != NULL) {
		DestroyImage(img);
	}
	if (info != NULL) {
		DestroyImageInfo(info);
	}
	if (sampled != NULL) {
		DestroyImage(sampled);
	}
	if (scaled != NULL) {
		DestroyImage(scaled);
	}
	if (err == 0) {
		return thumb->buf == NULL;
	}
	return err;
}

// Convert thumbnail to apropriate file type and write to buffer
static int writeThumb(Image *img, struct Thumbnail *thumb, ExceptionInfo *ex)
{
	ImageInfo *info = CloneImageInfo(NULL);
	char *format = NULL;
	bool needPNG = false;

	if (strcmp(img->magick, "JPEG")) {
		int err = hasTransparency(img, &needPNG, ex);
		if (err) {
			DestroyImageInfo(info);
			return err;
		}
	}
	if (needPNG) {
		format = "PNG";
		info->quality = 75;
		thumb->isPNG = true;
	} else {
		format = "JPEG";
		info->quality = 90;
	}
	strcpy(info->magick, format);
	strcpy(img->magick, format);
	thumb->buf = ImageToBlob(info, img, &thumb->size, ex);

	DestroyImageInfo(info);
	return 0;
}

// Itterates over all pixels and checks, if any transparency present
static int hasTransparency(const Image *img, bool *needPNG, ExceptionInfo *ex)
{
	// No alpha channel
	if (!img->matte) {
		return 0;
	}

	// Transparent pixels are most likely to also be in the first row, so
	// retrieve one row at a time. It is also more performant to retrieve entire
	// rows instead of individual pixels.
	for (unsigned long i = 0; i < img->rows; i++) {
		const PixelPacket *packets =
			AcquireImagePixels(img, 0, i, img->columns, 1, ex);
		if (packets == NULL) {
			return 1;
		}
		for (unsigned long j = 0; j < img->columns; j++) {
			if (packets[i].opacity != MaxRGB) {
				*needPNG = true;
				return 0;
			}
		}
	}

	return 0;
}
