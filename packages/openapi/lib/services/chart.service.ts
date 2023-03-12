import Chart from "../models/chart.model";


export async function getChart(limit) {
    return await Chart.find().sort({'created': -1}).limit(limit).exec()
}