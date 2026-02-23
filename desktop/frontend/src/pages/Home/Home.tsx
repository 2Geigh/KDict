import MainInput from "../../components/MainInput/MainInput"
import "./Home.scss"

const Home = () => {
	return (
		<div id="Home">
			<div id="container">
				<MainInput />
				<a id="viewCollection" href="">
					View collection
				</a>
			</div>
		</div>
	)
}

export default Home
