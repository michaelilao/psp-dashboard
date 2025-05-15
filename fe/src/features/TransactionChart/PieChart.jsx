import { PieChart, colors } from "../../components/PieChart/PieChart";

function TransactionPieChart({ transactions, title }) {
	const amountByCategory = transactions.reduce((accum, curr) => {
		if (curr.category in accum) {
			accum[curr.category] += curr.amount;
		} else {
			accum[curr.category] = curr.amount;
		}
		return accum;
	}, {});

	const labels = [];
	const values = [];
	Object.entries(amountByCategory).forEach(([cat, amount]) => {
		labels.push(cat);
		values.push(amount);
	});

	const chartData = {
		labels: labels,
		datasets: [
			{
				label: "Amount in $CAD",
				data: values,
				backgroundColor: colors,
				borderColor: colors,
				borderWidth: 1,
			},
		],
	};

	return <PieChart data={chartData} title={title} />;
}

export { TransactionPieChart };
