import { FC } from "react"
import { Link } from "react-router-dom"
import "./Navbar.scss"

const Navbar: FC = () => {
	return (
		<nav id="Navbar">
			<Link id="logo" to="/home">
				Vedana
			</Link>
			<ul>
				<li>
					<Link to={"/collection"}>Browse</Link>
				</li>
				<li>
					<Link to={"/stats"}>Stats</Link>
				</li>
				<li>
					<Link to={"/settings"}>Settings</Link>
				</li>
			</ul>
		</nav>
	)
}

export default Navbar
