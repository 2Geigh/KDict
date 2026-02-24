import { FC } from "react"
import "./Collection.scss"
import Sentence from "../../components/models/Sentence"

const Collection: FC = () => {
	// get rows data
	const sentences: Sentence[] = []
	const Rows = sentences.map((sentence, index) => (
		<tr>
			<td>
				<input
					type="checkbox"
					name="sentenceSelect"
					id={`check-${sentence.id}`}
					value={sentence.id}
				/>
			</td>
			<td>{sentence.id}</td>
			<td>{sentence.text}</td>
			<td>{sentence.target}</td>
			<td>{sentence.target_meaning}</td>
		</tr>
	))

	return (
		<div id="Collection" className="page">
			<div id="container">
				<div className="top">
					<select defaultValue={"sentences"}>
						<option value="sentences">Sentences</option>
						<option value="words">Words</option>
					</select>
					<button>Export</button>
				</div>
				<table>
					<tr>
						<th>
							<input
								type="checkbox"
								name="sentenceSelect"
								id="selectAll"
								value="selectAll"
							/>
						</th>
						<th>Id</th>
						<th>Sentence</th>
						<th>Target</th>
						<th>Target Meaning</th>
					</tr>
					{Rows}
				</table>
			</div>
		</div>
	)
}

export default Collection
