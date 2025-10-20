import sys
import openkorpos_dic
from konlpy.tag import Komoran
from konlpy.utils import pprint

# Initialize the Komoran object
komoran = Komoran()

def filterJunkWords(posOutput: list[tuple[str, str]]):
    toOutput = []
    for wordTuple in posOutput:
        # Filter out junk words based on specified tags
        if wordTuple[1] not in ["SF", "SY", "SC"]:
            toOutput.append(wordTuple)
    return toOutput

def parseSentence(query: str):
    # Parse the input query with Komoran
    unfilteredOutput = komoran.pos(query)

    # Filter the parsed output
    filteredOutput = filterJunkWords(unfilteredOutput)

    return filteredOutput

if __name__ == "__main__":

    numberOfArgumentsProvided = len(sys.argv)

    if numberOfArgumentsProvided > 2:
        print("Too many arguments provided.")
        sys.exit(1)

    if numberOfArgumentsProvided > 1:
        # Use repr to safely handle the input string
        argument = repr(sys.argv[1])[1:-1]  # Remove the quotes added by repr

        result = parseSentence(argument)
        
        for wordTuple in result:
            word = wordTuple[0]
            print(word)
        sys.exit(0)

    else:
        print("No input string provided.")
        sys.exit(-1)  # Exit with error code if no input string is provided
