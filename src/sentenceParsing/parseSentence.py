import timeit
from konlpy.tag import Mecab
from konlpy.utils import pprint
import openkorpos_dic

mecab = Mecab(dicpath=openkorpos_dic.DICDIR)

def filterJunkWords(mecabOutput: [tuple[str, str]]):
    toOutput = []
    for wordTuple in mecabOutput:
        if wordTuple[1] not in ["SF", "SY", "SC"]:
            toOutput.append(wordTuple)
    return toOutput

def parseSentence(query: [str]):
    unfilteredOutput = mecab.pos(query)
    filteredOutput = filterJunkWords(unfilteredOutput)
    # print(filteredOutput)
    return filteredOutput

if __name__ == "__main__":
    mecab_time = timeit.timeit('parseSentence("안녕하세요! 오늘은 2025년 10월 18일입니다. #KoreanText @User123 이메일: example@test.com 웹사이트: http://example.com")',
                                setup="from __main__ import pprint, mecab, openkorpos_dic, parseSentence, filterJunkWords",
                                number=1000)

    # Print timing results
    print(f'Mecab time: {mecab_time}')
