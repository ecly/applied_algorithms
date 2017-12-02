for lang in cs de en fr hi; do 
    ./tnt/tnt models/${lang} data/${lang}.test > data/${lang}.test.tagged
done
