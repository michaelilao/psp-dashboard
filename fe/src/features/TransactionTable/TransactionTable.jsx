import {
	TrashIcon,
	PencilSquareIcon,
	PlusIcon,
} from "@heroicons/react/24/outline";
import { Button } from "../../components/Button/Button";
import { formatCurrency, formatDate } from "../../utils/utils";
import { useState } from "react";
import { DeleteTransactionModal } from "../TransactionModal/DeleteModal";
import { CreateTransactionModal } from "../TransactionModal/CreateModal";

const COLS = [
	{ id: "name", label: "Name" },
	{ id: "amount", label: "Amount", getFormat: (num) => formatCurrency(num) },
	{ id: "category", label: "Category" },
	{ id: "transactionType", label: "Type" },
	{ id: "date", label: "Date", getFormat: (s) => formatDate(s) },
	{ id: "notes", label: "Notes" },
];

function TransactionTable({ transactions, userId }) {
	const [modifyTransaction, setModifyTransaction] = useState({
		type: "",
		transaction: null,
	});

	const getModal = () => {
		if (modifyTransaction.type == "") {
			return;
		}

		const onClose = () => {
			setModifyTransaction({
				type: "",
				transaction: null,
			});
		};

		if (modifyTransaction.type == "new") {
			return <CreateTransactionModal onClose={onClose} userId={userId} />;
		}

		if (modifyTransaction.type == "delete") {
			return (
				<DeleteTransactionModal
					transaction={modifyTransaction.transaction}
					onClose={onClose}
				/>
			);
		}
	};

	return (
		<>
			{getModal()}
			<div className="overflow-x-auto">
				<div className="flex align-baseline gap-4">
					<h2 className="text-2xl font-bold text-blue-500">Transactions</h2>
					<Button
						isIcon
						onClick={() => {
							setModifyTransaction({
								transaction: {},
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
														transaction: t,
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
														transaction: t,
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
