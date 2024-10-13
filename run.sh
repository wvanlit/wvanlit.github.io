rm -rf ./bin ./out

go build -C ./generator -o ../bin/site-generator

./bin/site-generator \
    --content content/ \
    --templates templates/ \
    --static static/ \
    --output out/ \
    $@