#include <vector>
#include "cv.h"

matptr_t nullPtr()
{
    return nullptr;
}

matptr_t cvNew()
{
    return new cv::Mat();
}

matptr_t cvNewFromBytes(int rows, int cols, int type, unsigned char *buf)
{
    return new cv::Mat(rows, cols, type, buf);
}

matptr_t cvIMRead(const char* filename, int flags)
{
    cv::Mat img = cv::imread(filename, flags);
    return new cv::Mat(img);
}

matptr_t cvIMDecode(unsigned char *buf, int length, int flags)
{
    std::vector<unsigned char> data(buf, buf + length);
    cv::Mat img = cv::imdecode(data, flags);
    return new cv::Mat(img);
}

void cvClose(matptr_t mat)
{
    if (!mat) return;
    delete (cv::Mat *)mat;
}

int cvIMWrite(const char* filename, const matptr_t img)
{
    if (!img) return 0;

    return cv::imwrite(filename, *img) ? 1 : 0;
}

int cvImageIMWriteWithParams(const char* filename, matptr_t img, int* params, int length)
{
    if (!img) return 0;

    std::vector<int> compression_params(params, params+length);

    return cv::imwrite(filename, *img, compression_params) ? 1 : 0;
}

void cvResize(matptr_t src, matptr_t dst, int dw, int dh, double fx, double fy, int interp)
{
    cv::Size sz(dw, dh);
    cv::resize(*src, *dst, sz, fx, fy, interp);
}

int cvEmpty(matptr_t mat)
{
    return mat->empty();
}

void cvtColor(matptr_t src, matptr_t dst, int code)
{
    cv::cvtColor(*src, *dst, code);
}


size_t cvSize(matptr_t m, int **ids)
{
    cv::MatSize ms(m->size);
    size_t length = ms.dims();
    *ids = (int *)malloc(length * sizeof(int));

    for (size_t i = 0; i < length; ++i) {
        (*ids)[i] = ms[i];
    }

    return length;
}
