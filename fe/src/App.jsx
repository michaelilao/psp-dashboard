import { Header } from "./components/Header/Header";
import { SideBar } from "./components/SideBar/SideBar";
import { Routes, Route } from "react-router";
import { Home } from "./pages/Home/Home";
import { User } from "./pages/User/User";

function App() {
	return (
		<div className="flex h-dvh">
			<div className="h-full w-full md:w-1/10">
				<SideBar />
			</div>
			<div className="bg-slate-50 w-full p-14">
				<Header />
				<Routes>
					<Route path="/" element={<Home />} />
					<Route path="/dashboard" element={<Home />} />
					<Route path="/users" element={<User />} />
					<Route path="/users/:userId" element={<User />} />
				</Routes>
			</div>
		</div>
	);
}

export default App;
