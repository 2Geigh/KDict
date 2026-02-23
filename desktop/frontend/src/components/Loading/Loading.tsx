import { FC } from "react"
import "./Loading.scss"

type LoadingProps = {
	text: string
}
const Loading: FC<LoadingProps> = ({ text }) => {
	return (
		<div id="Loading">
			<img
				src="/stroke-order/fast/fast_直-order.gif"
				alt={`${text}...`}
				id="loadingIcon"
			/>
			<span>{text}</span>
		</div>
	)
}

export default Loading
