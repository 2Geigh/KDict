import { Link } from "react-router-dom"
import MainInput from "../../components/MainInput/MainInput"
import "./Home.scss"

const Home = () => {
	return (
		<div id="Home" className="page">
			<div id="container">
				<MainInput />
			</div>
		</div>
	)
}

export default Home
