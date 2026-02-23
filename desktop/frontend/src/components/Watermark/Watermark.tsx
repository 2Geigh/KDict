import { FC } from "react"
import "./Watermark.scss"

export const Watermark: FC = () => {
	const thisYear = new Date().getFullYear()

	return (
		<div id="Watermark">
			<a href="https://nicholasgarcia.com" target="_blank">
				Nicholas Garcia
			</a>
			,
			<span className="year">
				{" "}
				{thisYear > 2026 ? `2026-${thisYear}` : thisYear}.
			</span>
		</div>
	)
}
