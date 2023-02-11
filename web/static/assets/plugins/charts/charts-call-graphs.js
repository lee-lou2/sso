//Color Variables
var redFull = '#BF263C';
var redFade = 'rgba(191, 38, 60, 0.2)';
var yellowFull = '#F6BB42';
var yellowFade = 'rgba(246, 187, 66, 0.2)';
var greenFull = '#8CC152';
var greenFade = 'rgba(140, 193, 82, 0.1)';
var greenFade2 = 'rgba(140, 193, 82, 0.7)';
var blueFull = '#5D9CEC';
var blueFade = 'rgba(93, 156, 236, 0.2)';
var blueFade2 = 'rgba(93, 156, 236, 0.7)';
var magentaFull = '#AC92EC';
var magentaFade = 'rgba(172, 146, 236, 0.2)';
var grayFull = '#CCD1D9';
var grayFade = 'rgba(204, 209, 217, 0.2)';
var orangeFull = '#E9573F';
var orangeFade = 'rgba(233, 87, 63, 0.2)'
var facebookColor = 'rgba(66,103,178, 0.8)';
var twitterColor = 'rgba(29,161,242, 0.8)';
var whatsappColor = 'rgba(37,211,102,0.8)';

//Component Graphs
//Component Graphs
//Component Graphs

var verticalChart = document.querySelectorAll('#vertical-chart')[0]
var horizontalChart = document.querySelectorAll('#horizontal-chart')[0]
var lineChart = document.querySelectorAll('#line-chart')[0]
if (verticalChart) {
	var verticalDemoChart = new Chart(verticalChart, {
		type: 'bar',
		data: {
			labels: ["2010", "2015", "2020", "2025"],
			datasets: [{
				label: "iOS",
				backgroundColor: greenFade2,
				data: [300, 600, 700, 900]
			}, {
				label: "Android",
				backgroundColor: blueFade2,
				data: [390, 550, 800, 1000]
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			title: {
				display: false
			}
		}
	});
}

if (horizontalChart) {
	var horizontalDemoChart = new Chart(horizontalChart, {
		type: 'bar',
		data: {
			labels: ["2010", "2013", "2016", "2020"],
			datasets: [{
				label: "Mobile",
				backgroundColor: yellowFull,
				data: [330, 400, 580, 590]
			}, {
				label: "Responsive",
				backgroundColor: orangeFull,
				data: [390, 450, 550, 570]
			}]
		},
		options: {
			indexAxis: 'y',
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			title: {
				display: false
			}
		}
	});
}

if (lineChart) {
	var lineDemoChart = new Chart(lineChart, {
		type: 'line',
		data: {
			labels: [2010, 2012, 2014, 2016, 2018, 2020, 2022],
			datasets: [{
				data: [77, 75, 70, 65, 63, 65, 70, 95],
				label: "Desktop Web",
				fill:true,
				backgroundColor:blueFade,
				borderColor: blueFull,
				lineTension: 0.3,
				pointRadius: 0,
			}, {
				data: [62, 64, 65, 66, 73, 75, 80, 95],
				label: "Mobile Web",
				fill:true,
				backgroundColor:greenFade,
				borderColor: greenFull,
				lineTension: 0.3,
				pointRadius: 0,
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			title: {
				display: false
			}
		}
	});
}

//Component Charts
//Component Charts
//Component Charts

var pieChart = document.querySelectorAll('#pie-chart')[0]
var doughnutChart = document.querySelectorAll('#doughnut-chart')[0]
var polarChart = document.querySelectorAll('#polar-chart')[0]
if (pieChart) {
	var pieDemoChart = new Chart(pieChart, {
		type: 'pie',
		data: {
			labels: ["Facebook", "Twitter", "WhatsApp"],
			datasets: [{
				backgroundColor: [facebookColor, twitterColor, whatsappColor],
				borderColor: "rgba(255,255,255,0.1)",
				data: [7000, 3000, 2000]
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			tooltips: {
				enabled: true
			},
			animation: {
				duration: 1500
			}
		}
	});
}

if (doughnutChart) {
	var doughnutDemoChart = new Chart(doughnutChart, {
		type: 'doughnut',
		data: {
			labels: ["Apple", "Samsung", "Google"],
			datasets: [{
				backgroundColor: [grayFull, blueFull, orangeFull],
				borderColor: "rgba(255,255,255,0.1)",
				data: [5500, 4000, 3000]
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			tooltips: {
				enabled: true
			},
			animation: {
				duration: 1500
			},
			layout: {
				padding: {
					bottom: 30
				}
			}
		}
	});
}

if (polarChart) {
	var polarDemoChart = new Chart(polarChart, {
		type: 'polarArea',
		data: {
			labels: ["Windows", "Mac", "Linux"],
			datasets: [{
				backgroundColor: [grayFull, blueFull, orangeFull],
				borderColor: "rgba(255,255,255,0.1)",
				data: [7000, 10000, 5000]
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			tooltips: {
				enabled: true
			},
			animation: {
				duration: 1500
			},
			layout: {
				padding: {
					bottom: 30
				}
			}
		}
	});
}

//Page Wallet
//Page Wallet
//Page Wallet

var walletChart = document.querySelectorAll('#wallet-chart')[0]
if (walletChart) {
	var walletDemoChart = new Chart(walletChart, {
		type: 'bar',
		data: {
			labels: ["Jun", "Jul", "Aug"],
			datasets: [{
				label: "Expenses",
				backgroundColor: "#ED5565",
				data: [100, 100, 130]
			}, {
				label: "Income",
				backgroundColor: "#A0D468",
				data: [170, 145, 165]
			}, {
				label: "Subscriptions",
				backgroundColor: "#4A89DC",
				data: [71, 50, 70]
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: true,
					position: 'bottom',
					labels: {
						fontSize: 13,
						padding: 15,
						boxWidth: 12
					},
				},
			},
			title: {
				display: false
			}
		}
	});
}


//Follower Stats
var followerStats = document.querySelectorAll('#followerStats')[0]
if (followerStats) {
	var followerStatsDemo = new Chart(followerStats, {
		type: 'line',
		data: {
			labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			datasets: [{
				lineTension: 0.3,
				label: "Followers",
				backgroundColor: blueFade,
				borderColor: blueFull,
				fill: true,
				borderWidth: 2,
				data: [110, 115, 108, 115, 118, 121, 122, 100]
			}]
		},
		options: {
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: false,
				},
			},
			title: {
				display: false
			}
		}
	});
}

//Income Stats
var incomeStats = document.querySelectorAll('#incomeStats')[0]
if (incomeStats) {
	var incomeStatsDemo = new Chart(incomeStats, {
		type: 'line',
		data: {
			labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			datasets: [{
				lineTension: 0.3,
				label: "Income Stats",
				backgroundColor: greenFade,
				borderColor:greenFull,
				fill: true,
				borderWidth: 2,
				data: [110, 115, 112, 115, 118, 121, 122, 100, 130],
				pointRadius: 0,
			}]
		},
		options: {
			scales: {
				yAxis: {display: false,},
				xAxis: {display: false,}
			},
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: false,
				},
			},
			title: {
				display: false
			}
		}
	});
}

//New User Stats
var newUserStats = document.querySelectorAll('#newUserStats')[0]
if (newUserStats) {
	var newUserStatsDemo = new Chart(newUserStats, {
		type: 'line',
		data: {
			labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			datasets: [{
				lineTension: 0.3,
				label: "New Users",
				backgroundColor: blueFade,
				borderColor: blueFull,
				fill: true,
				borderWidth: 2,
				data: [120, 115, 112, 115, 118, 118, 116, 100, 130],
				pointRadius: 0,
			}]
		},
		options: {
			scales: {
				yAxis: {display: false,},
				xAxis: {display: false,}
			},
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: false,
				},
			},
			title: {
				display: false
			}
		}
	});
}

//User Reach Stats
var userReachStats = document.querySelectorAll('#userReachStats')[0]
if (userReachStats) {
	var userReachStatsDemo = new Chart(userReachStats, {
		type: 'line',
		data: {
			labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			datasets: [{
				lineTension: 0.3,
				label: "User Reach",
				backgroundColor: yellowFade,
				borderColor: yellowFull,
				fill: true,
				borderWidth: 2,
				data: [120, 115, 112, 115, 118, 118, 116, 100, 130],
				pointRadius: 0,
			}]
		},
		options: {
			scales: {
				yAxis: {display: false,},
				xAxis: {display: false,}
			},
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: false,
				},
			},
			title: {
				display: false
			}
		}
	});
}

//User Interaction Stats
var userInteractionStats = document.querySelectorAll('#userInteractionStats')[0]
if (userInteractionStats) {
	var userInteractionStatsDemo = new Chart(userInteractionStats, {
		type: 'line',
		data: {
			labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			datasets: [{
				lineTension: 0.3,
				label: "User Interactions",
				backgroundColor: magentaFade,
				borderColor: magentaFull,
				fill: true,
				borderWidth: 2,
				data: [110, 113, 112, 114, 115, 115, 119, 100],
				pointRadius: 0,
			}]
		},
		options: {
			scales: {
				yAxis: {display: false,},
				xAxis: {display: false,}
			},
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: false,
				},
			},
			title: {
				display: false
			}
		}
	});
}