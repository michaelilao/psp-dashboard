import { Header } from "./components/Header/Header";
import { SideBar } from "./components/SideBar/SideBar";
import { Routes, Route } from "react-router";
import { Home } from "./pages/Home/Home";
import { Users } from "./pages/Users/User";
import { Transactions } from "./pages/Transactions/Transactions";

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
					<Route path="/users" element={<Users />} />
					<Route path="/transactions" element={<Transactions />} />
				</Routes>
			</div>
		</div>
	);
}

export default App;
