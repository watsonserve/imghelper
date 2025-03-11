#ifndef _CR2_THUNBNAIL_H_
#define _CR2_THUNBNAIL_H_

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/sendfile.h>

#define tImageWidth                256
#define tImageLength               257
#define tBitsPerSample             258
#define tCompression               259
#define tPhotometricInterpretation 262

#define tStripOffsets    273
#define tSamplesPerPixel 277
#define tRowsPerStrip    278
#define tStripByteCounts 279

#define tTileWidth      322
#define tTileLength     323
#define tTileOffsets    324
#define tTileByteCounts 325

#define tXResolution    282
#define tYResolution    283
#define tResolutionUnit 296

#define tPredictor    317
#define tColorMap     320
#define tExtraSamples 338
#define tSampleFormat 339

struct ifd_t
{
    uint16_t tag;
    uint16_t type;
    uint32_t count;
    uint32_t data;
};

struct cr2_head_t
{
    int16_t byteOrder;
    int16_t sign;
    int32_t offset;
};

struct meta_t
{
    uint16_t formatType;
    uint16_t width;
    uint16_t height;
    uint32_t stripOffset;
    uint32_t stripLen;
};

void cr2_thumbnail(int fd, off_t *off, size_t *len);

#endif
