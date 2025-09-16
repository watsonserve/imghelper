#include <vector>
#include "cv.h"

Mat cvNew()
{
    return new cv::Mat();
}

Mat cvNewFromBytes(int rows, int cols, int type, unsigned char *buf)
{
    return new cv::Mat(rows, cols, type, buf);
}

Mat cvIMRead(const char* filename, int flags)
{
    cv::Mat img = cv::imread(filename, flags);
    return new cv::Mat(img);
}

Mat cvIMDecode(unsigned char *buf, int length, int flags)
{
    std::vector<unsigned char> data(buf, buf + length);
    cv::Mat img = cv::imdecode(data, flags);
    return new cv::Mat(img);
}

void cvClose(Mat mat)
{
    if (!mat) return;
    delete (cv::Mat *)mat;
}

int cvIMWrite(const char* filename, const Mat img)
{
    if (!img) return 0;

    return cv::imwrite(filename, *img) ? 1 : 0;
}

int cvImageIMWriteWithParams(const char* filename, Mat img, int* params, int length)
{
    if (!img) return 0;

    std::vector<int> compression_params(params, params+length);

    return cv::imwrite(filename, *img, compression_params) ? 1 : 0;
}

void cvResize(Mat src, Mat dst, int dw, int dh, double fx, double fy, int interp)
{
    cv::Size sz(dw, dh);
    cv::resize(*src, *dst, sz, fx, fy, interp);
}

int cvEmpty(Mat mat)
{
    return mat->empty();
}

void cvtColor(Mat src, Mat dst, int code)
{
    cv::cvtColor(*src, *dst, code);
}


size_t cvSize(Mat m, int **ids)
{
    cv::MatSize ms(m->size);
    size_t length = ms.dims();
    *ids = (int *)malloc(length * sizeof(int));

    for (size_t i = 0; i < length; ++i) {
        (*ids)[i] = ms[i];
    }

    return length;
}
