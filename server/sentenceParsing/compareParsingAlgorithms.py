import timeit
from konlpy.tag import Hannanum, Kkma, Komoran, Mecab, Okt
from konlpy.utils import pprint
import openkorpos_dic

# Initialize taggers
kkma = Kkma()
hannanum = Hannanum()
komoran = Komoran()
mecab = Mecab(dicpath=openkorpos_dic.DICDIR)
okt = Okt()

# Function to time the POS tagging
def test_pos_tagging(phrase: str, replicate: int):
    # Use timeit to measure execution time of each tagger
    kkma_time = timeit.timeit(
        f"pprint(kkma.pos('''{phrase}'''))",  # Use triple quotes for the phrase
        setup="from __main__ import pprint, kkma", 
        number=replicate
    )

    hannanum_time = timeit.timeit(
        f"pprint(hannanum.pos('''{phrase}'''))",
        setup="from __main__ import pprint, hannanum", 
        number=replicate
    )

    komoran_time = timeit.timeit(
        f"pprint(komoran.pos('''{phrase}'''))", 
        setup="from __main__ import pprint, komoran", 
        number=replicate
    )

    mecab_time = timeit.timeit(
        f"pprint(mecab.pos('''{phrase}'''))", 
        setup="from __main__ import pprint, mecab, openkorpos_dic", 
        number=replicate
    )

    okt_time = timeit.timeit(
        f"pprint(okt.pos('''{phrase}'''))", 
        setup="from __main__ import pprint, okt", 
        number=replicate
    )

    # Print timing results
    print(f'Phrase: {phrase}')
    print(f'Kkma time: {kkma_time}')
    print(f'Hannanum time: {hannanum_time}')
    print(f'Komoran time: {komoran_time}')
    print(f'Mecab time: {mecab_time}')
    print(f'OkT time: {okt_time}\n')

# Test phrases
test_phrases = {
    "long sentence": u'국제연합의 모든 사람들은 그 헌장에서 기본적 인권입니다.',
    "english letters": u'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
    "numbers and symbols": u'0123456789!@#$%^&*()_+-=[]{}|;:\'",.<>?',
    "greek letters": u'αβγδεζηθικλμνξοπρστυφχψωΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ',
    "cyrillic letters": u'йи́йло́гика́бро́шьЛе́вДави́д',
    "chinese characters": u'我好你好吗',
    "hebrew characters": u'שלום, מה שלומך?',
    "emoji and symbols": u'😊💡🚀✨🔍✈️🌍'
}

# Test each phrase
for key in test_phrases:
    test_pos_tagging(test_phrases[key], 1)
