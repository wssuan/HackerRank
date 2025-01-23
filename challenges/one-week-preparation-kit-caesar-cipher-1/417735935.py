#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'caesarCipher' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. STRING s
#  2. INTEGER k
#

def caesarCipher(s, k):
    ciphers = [chr(c) for c in range(256)]
    k %= 26
    d = k
    azaz = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
    AZAZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
    a, A = ord('a'), ord('A')
    for i in range(26):
        ciphers[a + i] = azaz[i + k]
        ciphers[A + i] = AZAZ[i + k]
    return "".join(ciphers[ord(c)] for c in s)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    s = input()

    k = int(input().strip())

    result = caesarCipher(s, k)

    fptr.write(result + '\n')

    fptr.close()
