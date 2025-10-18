# import konlpy
import timeit
from konlpy.tag import Hannanum
from konlpy.tag import Kkma
from konlpy.tag import Komoran
from konlpy.utils import pprint
from konlpy.tag import Mecab
import openkorpos_dic

kkma = Kkma()
hannanum = Hannanum()
komoran = Komoran()
mecab = Mecab(dicpath=openkorpos_dic.DICDIR)

# Timing with setup to import necessary variables
kkma_time = timeit.timeit("pprint(kkma.pos(u'국제연합의 모든 사람들은 그 헌장에서 기본적 인권, 인간의 존엄과 가치, 그리고 남녀의 동등한 권리에 대한 신념을 재확인하였으며, 보다 폭넓은 자유속에서 사회적 진보와 보다 나은 생활수준을 증진하기로 다짐하였고,'))", 
                           setup="from __main__ import pprint, kkma", 
                           number=10)

print("\n")
hannanum_time = timeit.timeit("pprint(hannanum.pos(u'국제연합의 모든 사람들은 그 헌장에서 기본적 인권, 인간의 존엄과 가치, 그리고 남녀의 동등한 권리에 대한 신념을 재확인하였으며, 보다 폭넓은 자유속에서 사회적 진보와 보다 나은 생활수준을 증진하기로 다짐하였고,'))", 
                               setup="from __main__ import pprint, hannanum", 
                               number=10)

print("\n")
komoran_time = timeit.timeit("pprint(komoran.pos(u'국제연합의 모든 사람들은 그 헌장에서 기본적 인권, 인간의 존엄과 가치, 그리고 남녀의 동등한 권리에 대한 신념을 재확인하였으며, 보다 폭넓은 자유속에서 사회적 진보와 보다 나은 생활수준을 증진하기로 다짐하였고,'))", 
                               setup="from __main__ import pprint, komoran", 
                               number=10)

print("\n")
mecab_time = timeit.timeit("pprint(mecab.pos(u'국제연합의 모든 사람들은 그 헌장에서 기본적 인권, 인간의 존엄과 가치, 그리고 남녀의 동등한 권리에 대한 신념을 재확인하였으며, 보다 폭넓은 자유속에서 사회적 진보와 보다 나은 생활수준을 증진하기로 다짐하였고,'))", 
                               setup="from __main__ import pprint, mecab, openkorpos_dic", 
                               number=10)

# Print timing results
print(f'Kkma time: {kkma_time}')
print(f'Hannanum time: {hannanum_time}')
print(f'Komoran time: {komoran_time}')
print(f'Mecab time: {mecab_time}')
