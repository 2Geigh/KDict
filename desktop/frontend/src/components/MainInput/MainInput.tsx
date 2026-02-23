import { ChangeEvent, FC, JSX, useEffect, useState } from "react"
import "./MainInput.scss"
import "../../styles/App.scss"
import Loading from "../Loading/Loading"
import { ParseSentence } from "../../util/language"
import { RandomHaiku } from "../../util/placeholder"

const MAX_SENTENCE_LENGTH: number = 500
const MAX_MEANING_LENGTH: number = 500

const MainInput: FC = () => {
	const [isParsing, setIsParsing] = useState<boolean>(false)
	const [sentence, setSentence] = useState<string>("")
	const [hasParsed, setHasParsed] = useState<boolean | undefined>(undefined)
	const [parsedWords, setParsedWords] = useState<string[] | undefined>(
		undefined,
	)
	const [targetWord, setTargetWord] = useState<string | undefined>(undefined)

	useEffect(() => {
		if (isParsing) {
			ParseSentence(sentence)
				.catch((err) => {
					setParsedWords(undefined)
					throw new Error(`couldn't parse sentence: ${err}`)
				})
				.then((words) => {
					setParsedWords(words)
				})
				.finally(() => {
					setIsParsing(false)
					setHasParsed(true)
				})
		}
	}, [isParsing])

	useEffect(() => {
		console.log(`sentence: ${sentence}`)
	}, [sentence])

	useEffect(() => {
		console.log("###### DEV_DIAGNOSTICS ######")
		console.log(`isParsing: ${isParsing}`)
		console.log(`hasParsed: ${hasParsed}`)
		console.log(`sentence: ${sentence}`)
		console.log(`parsedWords: ${parsedWords}`)
		console.log(`targetWord: ${targetWord}`)
		console.log("#############################")
	}, [isParsing, hasParsed, sentence, parsedWords, targetWord])

	let words: JSX.Element[] = []
	if (parsedWords) {
		words = parsedWords!.map((word, index) => (
			<button
				key={`${index}-${word}`}
				className="word"
				onClick={(e) => {
					const selected = e.currentTarget.children[1].textContent
					if (selected) {
						setTargetWord(selected)
					} else {
						setTargetWord(undefined)
					}
				}}
			>
				<span className="number_key">( {index} )</span>
				<li className="text">{word}</li>
			</button>
		))
	}

	return (
		<div id="MainInput">
			<input
				id="mainInput"
				placeholder={"Type a word or sentence..."}
				maxLength={MAX_SENTENCE_LENGTH}
				onChange={(e: ChangeEvent<HTMLInputElement>) => {
					e.preventDefault
					const input = e.target.value.trim()

					setSentence(input)

					if (input) {
						setIsParsing(true)
						setHasParsed(false)
					} else if (!input) {
						setIsParsing(false)
						setHasParsed(undefined)
						setTargetWord(undefined)
					}
				}}
			/>
			<div
				id="underMainInput"
				className={
					(isParsing || hasParsed) && sentence.trim().length > 0 ?
						"visible"
					:	"hidden"
				}
			>
				{isParsing && <Loading text="Loading..." />}
				{hasParsed && (
					<>
						<div id="selection">
							<div className="section word_selection">
								<span className="prompt select_a_target_word">
									Select a target word:
								</span>
								<ul className="words">{words}</ul>
							</div>
							<div className="section definition">
								<span className="prompt write_your_comprehension_notes">
									{targetWord ?
										`${targetWord}'s meaning`
									:	`Meaning`}
									:
								</span>
								<textarea
									name="meaning"
									id="meaning"
									placeholder={RandomHaiku()}
									maxLength={MAX_MEANING_LENGTH}
								></textarea>
							</div>
						</div>
						<button className="create_note">Create note</button>
					</>
				)}
			</div>
		</div>
	)
}

export default MainInput
