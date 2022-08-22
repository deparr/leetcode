#include <stdio.h>

bool power_of_four(int n) {
    return n > 0 && (n & (n-1)) == 0 && n & 0x55555555;
}

int main() {
    printf("%d", power_of_four(16));
}