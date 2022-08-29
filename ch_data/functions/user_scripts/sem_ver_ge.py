#!/usr/bin/python3

import sys

from packaging import version

if __name__ == "__main__":
    for line in sys.stdin:
        args = line.split('\t')
        result = version.parse(args[0]) >= version.parse(args[1])
        print(int(result))
        sys.stdout.flush()
