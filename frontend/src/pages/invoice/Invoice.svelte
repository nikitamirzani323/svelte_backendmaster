<script>
    import Home from "../invoice/Home.svelte";
    export let path_api = ""
    export let font_size = "";
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master")
    let akses_page = false;
    async function initapp() {
        const res = await fetch(path_api+"api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "INVOICE_HOME",
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
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch(path_api+"api/invoice", {
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
            let no = 0
            let status_class = "";
            let winlose_class = "";
            let totalpasaran_class = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if (record[i]["invoice_status"] == "COMPLETED") {
                        status_class = "bg-[#ebfbee] text-[#6ec07b]"
                    } else {
                        status_class = "bg-[#fde3e3] text-[#ea7779]"
                    }
                    if (record[i]["invoice_winlose"] > 0) {
                        winlose_class = "text-blue-700"
                    } else {
                        winlose_class = "text-red-700"
                    }
                    if (record[i]["invoice_totalpasaran"] > 0) {
                        totalpasaran_class = "text-blue-700"
                    } else {
                        totalpasaran_class = "text-red-700"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["invoice_id"],
                            home_company: record[i]["invoice_company"],
                            home_create: record[i]["invoice_date"],
                            home_name: record[i]["invoice_name"],
                            home_winlose: record[i]["invoice_winlose"],
                            home_winloseclass: winlose_class,
                            home_totalpasaran: record[i]["invoice_totalpasaran"],
                            home_totalpasaranclass: totalpasaran_class,
                            home_status: record[i]["invoice_status"],
                            home_statusclass: status_class,
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 1000);
    };
    const handleLogout = (e) => {
        logout()
    }
    initapp();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {path_api}
        {font_size}
        {master}
        {token}
        {listHome}
        {totalrecord}/>
{/if}