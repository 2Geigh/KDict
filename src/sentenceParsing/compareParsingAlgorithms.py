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
    kkma_time = timeit.timeit(f"pprint(kkma.pos('{phrase}'))", 
                               setup="from __main__ import pprint, kkma", 
                               number=replicate)

    hannanum_time = timeit.timeit(f"pprint(hannanum.pos('{phrase}'))", 
                                   setup="from __main__ import pprint, hannanum", 
                                   number=replicate)

    komoran_time = timeit.timeit(f"pprint(komoran.pos('{phrase}'))", 
                                  setup="from __main__ import pprint, komoran", 
                                  number=replicate)

    mecab_time = timeit.timeit(f"pprint(mecab.pos('{phrase}'))", 
                                setup="from __main__ import pprint, mecab, openkorpos_dic", 
                                number=replicate)

    okt_time = timeit.timeit(f"pprint(okt.pos('{phrase}'))", 
                              setup="from __main__ import pprint, okt", 
                              number=replicate)

    # Print timing results
    print(f'Phrase: {phrase}')
    print(f'Kkma time: {kkma_time}')
    print(f'Hannanum time: {hannanum_time}')
    print(f'Komoran time: {komoran_time}')
    print(f'Mecab time: {mecab_time}')
    print(f'OkT time: {okt_time}')
    print("\n")

# Test phrases
test_phrases = [
    # u'국제연합의 모든 사람들은 그 헌장에서 기본적 인권, 인간의 존엄과 가치, 그리고 남녀의 동등한 권리에 대한 신념을 재확인하였으며, 보다 폭넓은 자유속에서 사회적 진보와 보다 나은 생활수준을 증진하기로 다짐하였고.',
    u'대한민국의 언어는 아름다운 언어입니다.',
    # u'자연은 언제나 우리에게 많은 교훈을 줍니다.'
]

# Test each phrase
for phrase in test_phrases:
    test_pos_tagging(phrase, 1)
