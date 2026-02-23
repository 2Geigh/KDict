import { FC } from "react"
import "./Loading.scss"

type LoadingProps = {
	text: string
}
const Loading: FC<LoadingProps> = ({ text }) => {
	return (
		<div id="Loading">
			<img
				src="../../../public/stroke-order/fast/fast_美-order-1.gif"
				alt="Parsing..."
				id="loadingIcon"
			/>
			<span>{text}</span>
		</div>
	)
}

export default Loading
