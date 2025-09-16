#ifndef _CV_HELPER_
#define _CV_HELPER_


#ifdef __cplusplus
#include <opencv2/opencv.hpp>

typedef cv::Mat* Mat;

extern "C" {
#else
typedef void* Mat;
#endif

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

Mat cvNew();

Mat cvNewFromBytes(int rows, int cols, int type, unsigned char *buf);

Mat cvIMRead(const char* filename, int flags);

Mat cvIMDecode(unsigned char* buf, int length, int flags);

void cvClose(Mat mat);

int cvEmpty(Mat mat);

int cvIMWrite(const char* filename, const Mat image);

int cvImageIMWriteWithParams(const char* filename, Mat img, int* params, int length);

void cvResize(Mat src, Mat dst, int dw, int dh, double fx, double fy, int interp);

void cvtColor(Mat src, Mat dst, int code);

size_t cvSize(Mat m, int **ids);

#ifdef __cplusplus
}
#endif

#endif
