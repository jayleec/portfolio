#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/portfolio" -importPath portfolio -srcPath "$SCRIPTPATH/src" -runMode dev
