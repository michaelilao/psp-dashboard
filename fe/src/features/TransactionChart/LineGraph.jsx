import { LineGraph } from "../../components/LineGraph/LineGraph";
import { PieChart, colors } from "../../components/PieChart/PieChart";
import { formatDate, formatDateYearMonth } from "../../utils/utils";

//yyyy-mm to May 2024
const shortDateToLong = (yyyyMm) => {
	const [year, month] = yyyyMm.split("-");
	const date = new Date(year, month);
	const formatter = new Intl.DateTimeFormat("en-US", {
		month: "long",
		year: "numeric",
	});

	return formatter.format(date);
};

function TransactionLineGraph({ transactions, title }) {
	const totalNetOverTime = transactions.reduce((accum, curr) => {
		let sign = 1;
		if (curr?.transactionType == "expense") {
			sign = -1;
		}
		const date = formatDateYearMonth(curr.date);
		if (date in accum) {
			accum[date] += curr.amount * sign;
		} else {
			accum[date] = curr.amount * sign;
		}
		return accum;
	}, {});

	const sortedByDate = Object.entries(totalNetOverTime).sort(
		(a, b) => a[0] > b[0]
	);

	const chartData = {
		labels: sortedByDate.map((item) => shortDateToLong(item[0])),
		datasets: [
			{
				label: "Net per Month in $CAD",
				data: sortedByDate.map((item) => item[1]),
				backgroundColor: ["rgba(255, 99, 132)"],
				borderColor: ["rgba(255, 99, 132)"],
				borderWidth: 1,
			},
		],
	};

	return <LineGraph title="Total Net Over Time in $CAD" data={chartData} />;
}

export { TransactionLineGraph };
