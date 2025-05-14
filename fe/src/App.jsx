import { Header } from "./components/Header/Header";
import { SideBar } from "./components/SideBar/SideBar";
import { Routes, Route } from "react-router";

function App() {
	return (
		<div className="flex h-dvh">
			<div className="h-full w-full md:w-1/10">
				<SideBar />
			</div>
			<div className="bg-slate-50 w-full">
				<Header />
				<Routes>
					<Route path="/" element={<div>Home</div>} />
					<Route path="/dashboard" element={<div>Dashboard</div>} />
					<Route path="/users" element={<div>Users</div>} />
					<Route path="/transactions" element={<div>Transactions</div>} />
				</Routes>
			</div>
		</div>
	);
}

export default App;
