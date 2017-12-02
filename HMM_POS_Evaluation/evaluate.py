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
                if (len(words) == 1): # assume irrelavant
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

def get_accuracy(s_data, g_data):
    """
    Caluclates POS tagging accuracy with respect to gold data.
    """
    total = 0
    correct = 0
    for s_sent, g_sent in zip(s_data, g_data):
        for (_, s_tag), (_, g_tag) in zip(s_sent, g_sent):
            total += 1
            if s_tag == g_tag:
                correct += 1
    print("total: {}, correct: {}".format(total, correct))
    return correct/total

accuracies = {lang:get_accuracy(data[lang][2], data[lang][3]) for lang in langs}
print (accuracies)

def get_accuracy_known_vs_unknown(s_data, g_data, lexicon):
    """
    Separates the accuracy for known and unknown words.
    Known = exists in the training set.
    """
    correct_known = 0
    correct_unknown = 0
    total_known = 0
    total_unknown = 0
    for s_sent, g_sent in zip(s_data, g_data):
        for (word, s_tag), (_, g_tag) in zip(s_sent, g_sent):
            if word in lexicon:
                total_known += 1
                correct_known += int(s_tag == g_tag)
            else:
                total_unknown += 1
                correct_unknown += int(s_tag == g_tag)
    return correct_known/total_known, correct_unknown/total_unknown

accuracies = {lang:get_accuracy(data[lang][2], data[lang][3]) for lang in langs}
print("known vs unknown")
print (accuracies)
