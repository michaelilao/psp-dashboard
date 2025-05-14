import {
	TrashIcon,
	PencilSquareIcon,
	PlusIcon,
} from "@heroicons/react/24/outline";
import { Button } from "../../components/Button/Button";
import { formatCurrency, formatDate } from "../../utils/utils";
import { useState } from "react";
import { Link } from "react-router";

const COLS = [
	{ id: "name", label: "Name" },
	{ id: "amount", label: "Amount", getFormat: (num) => formatCurrency(num) },
	{ id: "category", label: "Category" },
	{ id: "transactionType", label: "Type" },
	{ id: "date", label: "Date", getFormat: (s) => formatDate(s) },
	{ id: "notes", label: "Notes" },
];

function TransactionTable({ transactions }) {
	const [modifyTransaction, setModifyTransaction] = useState({
		type: "",
		transaction: null,
	});

	return (
		<>
			<div className="overflow-x-auto">
				<div className="flex align-baseline gap-4">
					<h2 className="text-2xl font-bold text-blue-500">Transactions</h2>
					<Button
						isIcon
						onClick={() => {
							setModifyTransaction({
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
						{transactions.map((t) => {
							return (
								<tr key={t.id}>
									{COLS.map((col) => {
										let value = t?.[col.id];
										let textColor = "text-gray-900";
										if ("getTextColor" in col) {
											textColor = col.getTextColor(value);
										}
										if ("getFormat" in col) {
											value = col.getFormat(value);
										}

										if (col.isLink) {
											return (
												<td
													className={`px-6 py-4 whitespace-nowrap text-sm ${textColor}`}
													key={col.id}
												>
													<Link
														to={`/users/${t.id}`}
														className="text-blue-500 hover:underline"
													>
														{value}
													</Link>
												</td>
											);
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
													setModifyTransaction({
														type: "edit",
														user: t,
													});
												}}
											>
												<PencilSquareIcon className="size-6 text-blue-500 hover:text-blue-900" />
											</Button>
											<Button
												isIcon
												onClick={() => {
													setModifyTransaction({
														type: "delete",
														user: t,
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

export { TransactionTable };
