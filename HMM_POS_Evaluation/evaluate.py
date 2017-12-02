# Mostly just contains Zeljko Agic's code from the Applied Algorithm course at ITU

import sys

def load_tagged_file(filename):
    """
    Reads the part-of-speech-tagged data in "word[\t]tag[\n]" format
    as a list of sentences, which are in turn lists of <word, tag> pairs.
    """
    unique_words = set()
    sentences = []
    current = []
    with open(filename) as file:
        for line in file:
            line = line. strip()
            if line:
                words = line.split("\t") # 0 = word, 1 = tag
                if (len(words) == 1):
                    continue
                current.append((words[0], words[1]))
                unique_words.add(words[0])
            else:  # empty line is sentence delimiter
                sentences.append(current)
                current = []
    return sentences, unique_words

langs = ["cs", "de", "en", "fr", "hi"]
data = {}
for lang in langs:
    training, lexicon = load_tagged_file("data/{}.train".format(lang))
    test, _ = load_tagged_file("data/{}.test".format(lang))
    tagged, _ = load_tagged_file("data/{}.test.tagged".format(lang))
    data[lang] = (lexicon, training, test, tagged)

print("lang\t#unique\t#train\t#test\t#tagged")

for lang, datasets in data.items():
    unique_train = len(datasets[0])
    words_train = sum([len(sentence) for sentence in datasets[1]])
    words_test = sum([len(sentence) for sentence in datasets[2]])
    words_tagged = sum([len(sentence) for sentence in datasets[3]])
    print("{:2}\t{}\t{}\t{}\t{}".format(lang, 
                                        unique_train, 
                                        words_train, 
                                        words_test, 
                                        words_tagged))
