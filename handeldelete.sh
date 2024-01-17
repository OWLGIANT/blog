#!/bin/bash
# 遍历目录中的文件根据target_directory  目录文件已有的删除当前文件夹匹配文件
echo "删除开始"
target_directory="CompressFiles"
for file in "$target_directory"/*; do
    if [ -f "$file" ]; then
      filename=$(basename "$file")
      patternFile="${filename%.gz}"
      rm  -f  "$patternFile"*
    fi
done
echo "删除结束"