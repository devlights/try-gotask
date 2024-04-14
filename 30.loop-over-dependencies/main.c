#include <stdio.h>

int main() {

#if defined(__x86_64__) || defined(_M_X64)
    printf("このプログラムは64ビットでコンパイルされました。\n");
#elif defined(__i386__) || defined(_M_IX86)
    printf("このプログラムは32ビットでコンパイルされました。\n");
#else
    printf("プラットフォームが判定できません。\n");
#endif

    return 0;
}