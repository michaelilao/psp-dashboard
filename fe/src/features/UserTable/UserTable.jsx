import {
	TrashIcon,
	PencilSquareIcon,
	PlusIcon,
} from "@heroicons/react/24/outline";
import { Button } from "../../components/Button/Button";
import { formatCurrency } from "../../utils/utils";
import { useState } from "react";
import { EditUserModal } from "../UserModal/EditModal";
import { DeleteUserModal } from "../UserModal/DeleteModal";
import { RegisterUserModal } from "../UserModal/RegisterModal";

const COLS = [
	{ id: "name", label: "Name" },
	{ id: "email", label: "Email" },
	{ id: "notes", label: "Notes" },
	{ id: "income", label: "Income", getFormat: (num) => formatCurrency(num) },
	{
		id: "expenses",
		label: "Expenses",
		getFormat: (num) => formatCurrency(num),
	},
	{
		id: "total",
		label: "Total",
		getFormat: (num) => formatCurrency(num),

		getTextColor: (num) => {
			if (num > 0) {
				return "text-green-600";
			} else if (num < 0) {
				return "text-red-500";
			}
			return "text-gray-900";
		},
	},
];

function UserTable({ users }) {
	const [modifyUser, setModifyUser] = useState({
		type: "",
		user: null,
	});

	const getModal = () => {
		if (modifyUser.type == "") {
			return;
		}
		if (modifyUser.type == "new") {
			return (
				<RegisterUserModal
					user={modifyUser.user}
					onClose={() => {
						setModifyUser({
							type: "",
							user: null,
						});
					}}
				/>
			);
		}

		if (modifyUser.type == "delete") {
			return (
				<DeleteUserModal
					user={modifyUser.user}
					onClose={() => {
						setModifyUser({
							type: "",
							user: null,
						});
					}}
				/>
			);
		}

		if (modifyUser.type == "edit") {
			return (
				<EditUserModal
					user={modifyUser.user}
					onClose={() => {
						setModifyUser({
							type: "",
							user: null,
						});
					}}
				/>
			);
		}
	};

	return (
		<>
			{getModal()}
			<div className="overflow-x-auto">
				<div className="flex align-baseline gap-4">
					<h2 className="text-2xl font-bold text-blue-500">Users</h2>
					<Button
						isIcon
						onClick={() => {
							setModifyUser({
								user: {},
								type: "new",
							});
						}}
					>
						<PlusIcon className="size-8 text-white rounded-md bg-blue-500 hover:bg-blue-900" />
					</Button>
				</div>
				<table className="divide-y divide-gray-200">
					<thead className="">
						<tr>
							{COLS.map((col) => {
								return (
									<th
										key={col.id}
										scope="col"
										className="px-6 py-3 text-left text-sm font-medium text-gray-700"
									>
										{col.label}
									</th>
								);
							})}
							<th
								scope="col"
								className="px-6 py-3 text-left text-sm font-medium text-gray-700"
							>
								Actions
							</th>
						</tr>
					</thead>
					<tbody className="divide-y divide-gray-200 ">
						{users.map((u) => {
							return (
								<tr key={u.id}>
									{COLS.map((col) => {
										let value = u?.[col.id];
										let textColor = "text-gray-900";
										if ("getTextColor" in col) {
											textColor = col.getTextColor(value);
										}
										if ("getFormat" in col) {
											value = col.getFormat(value);
										}

										return (
											<td
												className={`px-6 py-4 whitespace-nowrap text-sm ${textColor}`}
												key={col.id}
											>
												{value}
											</td>
										);
									})}
									<td>
										<div className="flex justify-center gap-2">
											<Button
												isIcon
												onClick={() => {
													setModifyUser({
														type: "edit",
														user: u,
													});
												}}
											>
												<PencilSquareIcon className="size-6 text-blue-500 hover:text-blue-900" />
											</Button>
											<Button
												isIcon
												onClick={() => {
													setModifyUser({
														type: "delete",
														user: u,
													});
												}}
											>
												<TrashIcon className="size-6 text-red-500 hover:text-red-900" />
											</Button>
										</div>
									</td>
								</tr>
							);
						})}
					</tbody>
				</table>
			</div>
		</>
	);
}

export { UserTable };
