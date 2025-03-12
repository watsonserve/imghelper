#include "cr2_thumbnail.h"

void cr2_thumbnail(int fd, off_t *off, size_t *len)
{
    static uint32_t typeLen[] = {0, 1, 1, 2, 4, 8, 1, 1, 2, 4, 8, 4, 8};
    struct cr2_head_t head;
    struct meta_t meta;
    short count;

    int ret = pread(fd, &head, sizeof(struct cr2_head_t), 0);
    ret = pread(fd, &count, 2, head.offset);
    struct ifd_t ifds[count];
    ret = pread(fd, ifds, count * 12, head.offset + 2);

    for (int i = 0; i < count; i++)
    {
        struct ifd_t *ifdp = ifds + i;
        uint16_t tag = ifdp->tag;
        uint16_t len = typeLen[ifdp->type];
        uint32_t val = ifdp->data;
        switch (tag)
        {
            case tCompression:
                meta.formatType = val;
                break;
            case tImageWidth:
                meta.width = val;
                break;
            case tImageLength:
                meta.height = val;
                break;
            case tStripOffsets:
                meta.stripOffset = val;
                break;
            case tStripByteCounts:
                meta.stripLen = val;
                break;
            default:
        }
    }
    *off = meta.stripOffset;
    *len = meta.stripLen;
}
