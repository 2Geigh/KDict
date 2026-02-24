import { Router, Routes, Route, BrowserRouter } from "react-router-dom"
import "./styles/App.scss"
import "@fontsource/inter"

import Home from "./pages/Home/Home"
import { Watermark } from "./components/Watermark/Watermark"
import Collection from "./pages/Collection/Collection"
import Navbar from "./components/Navbar/Navbar"

function App() {
	return (
		<>
			<BrowserRouter>
				<Navbar />

				<Routes>
					<Route path="/" element={<Home />} />
					<Route path="/home" element={<Home />} />
					<Route path="/collection" element={<Collection />} />
				</Routes>
			</BrowserRouter>

			<Watermark />
		</>
	)
}

export default App
