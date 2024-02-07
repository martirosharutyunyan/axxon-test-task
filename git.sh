git add .
git status
read -r -p "Enter commit message: " MESSAGE
git commit -m "${MESSAGE}"
git push
