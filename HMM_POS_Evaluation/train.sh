for lang in cs de en fr hi; do 
    ./tnt/tnt-para -o models/${lang} data/${lang}.train
done
