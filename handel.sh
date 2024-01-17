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

# 获取当前文件夹路径
current_folder=$(pwd)
# 计算文件数
file_count=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" | wc -l)
# 计算总资源大小
total_size=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" -exec du -csh {} + | tail -n 1 | awk '{print $1}')
# 打印结果
echo "文件数: $file_count"
echo "总资源大小: $total_size"

# 获取今天的日期的起始时间戳
today_start=$(date -d "$(date +"%Y-%m-%d") 00:00:00" +%s)

# 遍历当前目录下的CSV文件
for file in *.csv; do
  # 获取文件的修改时间（mtime）
  mtime=$(stat -c %Y "$file")

  # 判断文件是否在今天更新
  if [ "$mtime" -ge "$today_start" ]; then
       echo "文件在今天更新过 $today_start  $mtime"
  else
       echo "文件不是今天更新的  $today_start  $mtime"
        # 检查是否存在对应的.gz文件
        if [ ! -f "${file}.gz" ]; then
          # 如果.gz文件不存在，则使用gzip进行压缩
          gzip -c "$file" > "$folder_path"/"${file}.gz"
        else
          mv ${file}.gz "$folder_path"/"${file}.gz"
        fi
  fi
done

# 获取当前文件夹路径
current_folder=$(pwd)
# 计算文件数
file_count=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" | wc -l)
# 计算总资源大小
total_size=$(find "$current_folder" -maxdepth 1 -type f -name "*.csv*" -exec du -csh {} + | tail -n 1 | awk '{print $1}')
# 打印结果
echo "文件数: $file_count"
echo "总资源大小: $total_size"

echo "压缩完成。"