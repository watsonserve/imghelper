#ifndef _CV_HELPER_
#define _CV_HELPER_


#ifdef __cplusplus
#include <opencv2/opencv.hpp>

typedef cv::Mat* matptr_t;

extern "C" {
#else
typedef void* matptr_t;
#endif

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

matptr_t nullPtr();

matptr_t cvNew();

matptr_t cvNewFromBytes(int rows, int cols, int type, unsigned char *buf);

matptr_t cvIMRead(const char* filename, int flags);

matptr_t cvIMDecode(unsigned char* buf, int length, int flags);

void cvClose(matptr_t mat);

int cvEmpty(matptr_t mat);

int cvIMWrite(const char* filename, const matptr_t image);

int cvImageIMWriteWithParams(const char* filename, matptr_t img, int* params, int length);

void cvResize(matptr_t src, matptr_t dst, int dw, int dh, double fx, double fy, int interp);

void cvtColor(matptr_t src, matptr_t dst, int code);

size_t cvSize(matptr_t m, int **ids);

#ifdef __cplusplus
}
#endif

#endif
