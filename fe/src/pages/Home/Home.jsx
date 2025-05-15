import { useEffect } from "react";
import { UserTable } from "../../features/UserTable/UserTable";
import { useUserStore } from "../../store/users";

function Home() {
	const { users, fetchUsers } = useUserStore((state) => state);

	useEffect(() => {
		fetchUsers();
	}, [fetchUsers]);

	return (
		<div>
			<h1 className="text-2xl mb-6 font-semibold">
				Expense Management Dashboard
			</h1>
			<UserTable users={users} />
		</div>
	);
}

export { Home };
