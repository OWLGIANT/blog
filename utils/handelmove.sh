# 配置目标文件夹
folder_path="CompressFiles"

if [ -d "$folder_path" ]; then
  echo "文件夹已经存在"
else
  mkdir -p "$folder_path"
  echo "文件夹已创建"
fi

echo "移动开始"

today_start=$(date -d "$(date +"%Y-%m-%d")" +%s)

for file in *.csv.gz; do

  mtime=$(stat -c %Y "$file")

  if [ "$mtime" -ge "$today_start" ]; then
        echo "文件在今天更新过 $today_start  $mtime"
  else
        echo "文件是之前更新的  $today_start  $mtime"
        mv "$file"  "$folder_path"/"$file"
  fi
done

echo "移动完成"