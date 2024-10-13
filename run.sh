echo "cleaning existing folders"

rm -rf ./bin ./out

echo "building go binary"

go build -C ./generator -o ../bin/site-generator

echo "generating site"

./bin/site-generator \
    --content content/ \
    --templates templates/ \
    --static static/ \
    --output out/ \
    $@