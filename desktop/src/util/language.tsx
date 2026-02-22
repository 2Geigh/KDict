import { sleep } from "./sleep"

export async function ParseSentence(sentence: string): Promise<string[]> {
	console.log(`Parsing ${sentence}...`)

	let words: string[] = ["다음", "은", "한국어", "신문", "의", "목록", "이다"]

	await sleep(3000) // Use the sleep function
	console.log("Sentence parsed.")
	return words
}
