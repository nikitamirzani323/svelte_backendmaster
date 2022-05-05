<script>
    export let path_api = ""
    export let font_size = "";
    import dayjs from "dayjs";
	let listwinlose = [];
	let listwinlose_company = [];
	let listcompanydata = [];
	let pasaran = [];
	let pasarantogel = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master")
    let akses_page = false;
	let winlose_year = dayjs().format("YYYY")
	let winlose_year_1 = parseInt(winlose_year)-1;
	let text_chart = "WINLOSE "+winlose_year;
	let select_year = winlose_year;
	let select_company = "";
    async function initapp() {
		const res = await fetch(path_api+"api/init", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				page: "DASHBOARD-VIEW",
			}),
		});
		const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            akses_page = true;
            initHome(winlose_year);
            listcompany();
        }
		
	}
	async function initHome(e) {
		listwinlose = [];
		pasaran = [];
        const res = await fetch(path_api+"api/dashboardwinlose", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                year: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (let i = 0; i < record.length; i++) {
					listwinlose = [
						...listwinlose,
						{
							pasaran_name: record[i]["dashboardwinlose_nmcompany"],
							pasaran_winlose: record[i]["dashboardwinlose_detail"],
						},
					];
				}
				for (let j = 0; j < listwinlose.length; j++) {
					let totaldata = listwinlose[j].pasaran_winlose.length;
					let temp = [];
					for (let x = 0; x < totaldata; x++) {
						temp.push(
							parseInt(listwinlose[j].pasaran_winlose[x]["dashboardwinlose_winlose"])/1000
						);
					}
					pasaran = [
						...pasaran,
						{
							name: listwinlose[j].pasaran_name,
							data: temp,
						},
					];
				}
            }
        }
		createChart()
    }
	async function winlosecompany(e,z) {
		listwinlose_company = [];
		pasarantogel = [];
        const res = await fetch(path_api+"api/dashboardwinlosecompany", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                year: e,
                company: z,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (let i = 0; i < record.length; i++) {
					listwinlose_company = [
						...listwinlose_company,
						{
							pasaran_name: record[i]["dashboardcompanypasaran_nmpasaran"],
							pasaran_winlose: record[i]["dashboardcompanypasaran_detail"],
						},
					];
				}
				for (let j = 0; j < listwinlose_company.length; j++) {
					let totaldata = listwinlose_company[j].pasaran_winlose.length;
					let temp = [];
					for (let x = 0; x < totaldata; x++) {
						temp.push(
							parseInt(listwinlose_company[j].pasaran_winlose[x]["dashboardcompanypasaran_winlose"])/1000
						);
						
					}
					pasarantogel = [
						...pasarantogel,
						{
							name: listwinlose_company[j].pasaran_name,
							data: temp,
						},
					];
				}
            }
			console.log(pasarantogel)
        }
		createChartWinlose()
    }
	async function listcompany() {
		listwinlose = [];
		pasaran = [];
        const res = await fetch(path_api+"api/dashboardlistcompany", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (let i = 0; i < record.length; i++) {
					listcompanydata = [
						...listcompanydata,
						{
							company_idcompany: record[i]["company_idcompany"],
							company_nmcompany: record[i]["company_nmcompany"],
						},
					];
				}
            }
        }
    }
	const handleSelect = (event) => {
		// alert(event.target.value)
        initHome(event.target.value);
		text_chart = "WINLOSE "+event.target.value;
    };
	const handleSelectCompany = (event) => {
		// alert(event.target.value)
        winlosecompany("2022",event.target.value);
		text_chart = "WINLOSE "+event.target.value;
    };
	async function createChart() {
		Highcharts.chart("container", {
			chart: {
				type: "column",
			},
			title: {
				text: "",
			},
			subtitle: {
				text: "Source: TOTO",
			},
			xAxis: {
				categories: [
					"Jan",
					"Feb",
					"Mar",
					"Apr",
					"May",
					"Jun",
					"Jul",
					"Aug",
					"Sep",
					"Oct",
					"Nov",
					"Dec",
				],
				crosshair: true,
			},
			yAxis: {
				min: 0,
				title: {
					text: "Rupiah",
				},
			},
			tooltip: {
				headerFormat:
					'<span style="font-size:10px">{point.key}</span><table>',
				pointFormat:
					'<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f}</b></td></tr>',
				footerFormat: "</table>",
				shared: true,
				useHTML: true,
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0,
				},
			},
			series: pasaran,
		});
	}
	async function createChartWinlose() {
		Highcharts.chart("container_winlose", {
			chart: {
				type: "column",
			},
			title: {
				text: "",
			},
			subtitle: {
				text: "Source: TOTO",
			},
			xAxis: {
				categories: [
					"Jan",
					"Feb",
					"Mar",
					"Apr",
					"May",
					"Jun",
					"Jul",
					"Aug",
					"Sep",
					"Oct",
					"Nov",
					"Dec",
				],
				crosshair: true,
			},
			yAxis: {
				min: 0,
				title: {
					text: "Rupiah",
				},
			},
			tooltip: {
				headerFormat:
					'<span style="font-size:10px">{point.key}</span><table>',
				pointFormat:
					'<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f}</b></td></tr>',
				footerFormat: "</table>",
				shared: true,
				useHTML: true,
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0,
				},
			},
			series: pasarantogel,
		});
	}
    async function logout() {
		localStorage.clear();
		window.location.href = "/";
	}
    initapp();
</script>
{#if akses_page}
	<div class="flex w-full gap-2 px-2">
		<div class="bg-white shadow-lg p-5 w-1/2">
			<div class="flex flex-col gap-2">
				<div class="flex">
					<h1 class="text-slate-600 font-bold text-sm lg:text-2xl uppercase w-full">{text_chart}</h1>
					<div class="hidden sm:flex md:flex justify-end w-28">
						<select
							on:change={handleSelect}
							bind:value="{select_year}" 
							class="select select-bordered w-full rounded-sm focus:ring-0 focus:outline-none">
							<option disabled selected value="">--Pilih Year--</option>
							<option value="{winlose_year}">{winlose_year}</option>
							<option value="{winlose_year_1}">{winlose_year_1}</option>
						</select>
					</div>
				</div>
				<div class="hidden  sm:inline w-full scrollbar-thin scrollbar-thumb-sky-200 h-[500px] ">
					<div class="w-full h-[500px]" id="container" />
				</div>
			</div>
		</div>

		<div class="bg-white shadow-lg p-5 w-1/2">
			<div class="flex flex-col gap-2">
				<div class="flex">
					<h1 class="text-slate-600 font-bold text-sm lg:text-2xl uppercase w-full">PASARAN : {text_chart}</h1>
					<div class="hidden sm:flex md:flex justify-end w-60">
						<select
							on:change={handleSelectCompany}
							bind:value="{select_company}" 
							class="select select-bordered w-full rounded-sm focus:ring-0 focus:outline-none">
							<option disabled selected value="">--Pilih Company--</option>
							{#each listcompanydata as rec}
								<option value="{rec.company_idcompany}">{rec.company_nmcompany}</option>
							{/each}
						</select>
					</div>
				</div>
				<div class="hidden  sm:inline w-full scrollbar-thin scrollbar-thumb-sky-200 h-[500px] ">
					<div class="w-full h-[500px]" id="container_winlose" />
				</div>
			</div>
		</div>
	</div>
{/if}