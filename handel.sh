#!/bin/bash
# 配置目标文件夹
folder_path="CompressFiles"

# 检查文件夹是否存在
if [ -d "$folder_path" ]; then
  echo "文件夹已经存在"
else
  # 如果文件夹不存在，则创建
  mkdir -p "$folder_path"
  echo "文件夹已创建"
fi

current_folder=$(pwd)
file_count=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" | wc -l)
total_size=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" -exec du -csh {} + | tail -n 1 | awk '{print $1}')

echo "文件数: $file_count"
echo "总资源大小: $total_size"

today_start=$(date -d "$(date +"%Y-%m-%d")" +%s)

for file in *.csv; do
  mtime=$(stat -c %Y "$file")
  if [ "$mtime" -ge "$today_start" ]; then
        echo "文件在今天更新过 $today_start  $mtime"
  else
        echo "文件不是今天更新的  $today_start  $mtime"
        if [ ! -f "${file}.gz" ]; then
          gzip -c "$file" > "$folder_path"/"${file}.gz"
        else
          mv ${file}.gz "$folder_path"/"${file}.gz"
        fi
  fi
done

current_folder=$(pwd)
file_count=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" | wc -l)
total_size=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" -exec du -csh {} + | tail -n 1 | awk '{print $1}')

echo "文件数: $file_count"
echo "总资源大小: $total_size"
echo "压缩完成。"