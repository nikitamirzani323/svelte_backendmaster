<script>
    import { createEventDispatcher } from "svelte";
    import Button_custom from '../../components/button_custom.svelte'
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte'

    export let path_api = "";
    export let font_size = "";
    export let token = "";
    export let idcompany = "";
    export let listMasterPasaran = [];
    let listPeriode = [];
    let listInvoiceMember = [];
    let listInvoiceMembertemp = [];
    let listInvoicelistbet = [];
    let listInvoiceGroupPermainan = [];
    let isModal_info_invoice = false
    let modal_infoinvoice_width = "max-w-xl"
    let select_periode = "";
    let select_pasaran = "";
    let totalrecord = 0;
    let totalwinlose = 0;
    let totalwinlose_css = "";
    let subtotal_win_css = "";
    let subtotal_winlose_css = "";
    let subtotal_win_css_temp = "";
    let subtotal_winlose_css_temp = "";
    let subtotal_bet = 0;
    let subtotal_bayar = 0;
    let subtotal_cancel = 0;
    let subtotal_win = 0;
    let subtotal_winlose = 0;
    let subtotal_bet_temp = 0;
    let subtotal_bayar_temp = 0;
    let subtotal_cancel_temp = 0;
    let subtotal_win_temp = 0;
    let subtotal_winlose_temp = 0;
    let buttonLoading_flag = false;
    let buttonLoading_class = "";
    let msg_error = "";
    let searchlistbet = "";
    let filterlistbet = [];
    let invoice_no = ""
    let dispatch = createEventDispatcher();
    async function generate() {
        listPeriode = [];
        totalwinlose = 0;
        totalwinlose_css = "text-blue-700";
        let flag = false;
        msg_error = "";
        
        if (select_periode == "") {
            flag = true;
            msg_error += "The Pilih Periode is required<br>";
        }
        if (select_pasaran == "") {
            flag = true;
            msg_error += "The Pilih Pasaran is required<br>";
        }
        if (flag == false) {
            buttonLoading_flag = true;
            const res = await fetch(path_api+"api/companylistkeluaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "COMPANY_HOME",
                    company: idcompany,
                    periode: select_periode,
                    pasaran: parseInt(select_pasaran),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                let record = json.record;
                totalwinlose = json.totalwinlose;
                if(totalwinlose > 0){
                    totalwinlose_css = "text-blue-700";
                }else{
                    totalwinlose_css = "text-red-500";
                }
                let periode_status_class = "";
                let css_totalmember = "";
                let css_totalbet = "";
                let css_totalbayar = "";
                let css_totalcancel = "";
                let css_winlose = "";
                let css_winlosetemp = "";
                let css_selisih = "";
                let css_revisi = "";
                if (record != null) {
                    totalrecord = record.length;
                    for (var i = 0; i < record.length; i++) {
                        let selisih = parseInt(record[i]["company_pasaran_winlose"]) - parseInt(record[i]["company_pasaran_winlosetemp"])
                        if(record[i]["company_pasaran_status"] == "DONE"){
                            periode_status_class = "bg-[#ebfbee] text-[#6ec07b]"
                        }else{
                            periode_status_class = "bg-[#FFEB3B] text-black "
                        }
                        if (parseInt(selisih) > 0) {
                            css_selisih = "text-blue-700";
                        } else {
                            css_selisih = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_totalmember"]) > 0) {
                            css_totalmember = "text-blue-700";
                        } else {
                            css_totalmember = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_totalbet"]) > 0) {
                            css_totalbet = "text-blue-700";
                        } else {
                            css_totalbet = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_totaloutstanding"]) > 0) {
                            css_totalbayar = "text-blue-700";
                        } else {
                            css_totalbayar = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_totalcancelbet"]) > 0) {
                            css_totalcancel = "text-blue-700";
                        } else {
                            css_totalcancel = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_winlose"]) > 0) {
                            css_winlose = "text-blue-700";
                        } else {
                            css_winlose = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_winlosetemp"]) > 0) {
                            css_winlosetemp = "text-blue-700";
                        } else {
                            css_winlosetemp = "text-red-500";
                        }
                        if (parseInt(record[i]["company_pasaran_revisi"]) > 0) {
                            css_revisi = "text-blue-700";
                        } else {
                            css_revisi = "text-red-500";
                        }
                        listPeriode = [
                            ...listPeriode,
                            {
                                no: record[i]["company_pasaran_no"],
                                company_pasaran_invoice: record[i]["company_pasaran_invoice"].toString(),
                                company_pasaran_idcompp: record[i]["company_pasaran_idcompp"],
                                company_pasaran_code: record[i]["company_pasaran_code"],
                                company_pasaran_periode: record[i]["company_pasaran_periode"],
                                company_pasaran_name: record[i]["company_pasaran_name"],
                                company_pasaran_keluaran: record[i]["company_pasaran_keluaran"],
                                company_pasaran_tanggal: record[i]["company_pasaran_tanggal"],
                                company_pasaran_totalmember:record[i]["company_pasaran_totalmember"],
                                company_pasaran_totalmember_css: css_totalmember,
                                company_pasaran_totalbet: record[i]["company_pasaran_totalbet"],
                                company_pasaran_totalbet_css: css_totalbet,
                                company_pasaran_totaloutstanding:record[i]["company_pasaran_totaloutstanding"],
                                company_pasaran_totaloutstanding_css: css_totalbayar,
                                company_pasaran_totalcancelbet:record[i]["company_pasaran_totalcancelbet"],
                                company_pasaran_totalcancelbet_css:css_totalcancel,
                                company_pasaran_msgrevisi: record[i]["company_pasaran_msgrevisi"],
                                company_pasaran_revisi: record[i]["company_pasaran_revisi"],
                                company_pasaran_revisi_css: css_revisi,
                                company_pasaran_winlose: record[i]["company_pasaran_winlose"],
                                company_pasaran_winlose_css: css_winlose,
                                company_pasaran_winlosetemp: record[i]["company_pasaran_winlosetemp"],
                                company_pasaran_winlosetemp_css: css_winlosetemp,
                                company_pasaran_selisih: selisih,
                                company_pasaran_selisih_css: css_selisih,
                                company_pasaran_status: record[i]["company_pasaran_status"],
                                company_pasaran_status_css: periode_status_class,
                            },
                        ];
                    }
                }
            }
            buttonLoading_flag = false;
        }else{
            if(msg_error != ""){
                let temp_msg = msg_error
                dispatch("handleCallNotif", {
                        temp_msg
                });
            }
        }
    }
    async function invoice_member(e) {
        listInvoiceMember = []
        subtotal_bet = 0;
        subtotal_bayar = 0;
        subtotal_cancel = 0;
        subtotal_win = 0;
        subtotal_winlose = 0;
        const res = await fetch(path_api+"api/companyinvoicemember", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let css_totalwin = "";
               
                for (var i = 0; i < record.length; i++) {
                    if (parseInt(record[i]["totalwin"]) > 0) {
                        css_totalwin = "text-blue-700";
                    } else {
                        css_totalwin = "text-red-500";
                    }
                    subtotal_bet = subtotal_bet + parseInt(record[i]["totalbet"])
                    subtotal_bayar = subtotal_bayar + parseInt(record[i]["totalbayar"])
                    subtotal_cancel = subtotal_cancel + parseInt(record[i]["totalcancelbet"])
                    subtotal_win = subtotal_win + parseInt(record[i]["totalwin"])
                    if(subtotal_win > 0){
                        subtotal_win_css = "text-red-500";
                    }else{
                        subtotal_win_css = "text-blue-700";
                    }
                    listInvoiceMember = [
                        ...listInvoiceMember,
                        {
                            member: record[i]["member"],
                            totalbet: record[i]["totalbet"],
                            totalbayar: record[i]["totalbayar"],
                            totalcancelbet: record[i]["totalcancelbet"],
                            totalwin: record[i]["totalwin"],
                            totalwin_css: css_totalwin,
                        },
                    ];
                }
                subtotal_winlose = parseInt(subtotal_bayar) - parseInt(subtotal_cancel) - parseInt(subtotal_win)
                if(subtotal_winlose > 0){
                    subtotal_winlose_css = "text-blue-700";
                }else{
                    subtotal_winlose_css = "text-red-500";
                }
            }
        } 
    }
    async function invoice_membertemp(e) {
        listInvoiceMembertemp = []
        subtotal_bet_temp = 0;
        subtotal_bayar_temp = 0;
        subtotal_cancel_temp = 0;
        subtotal_win_temp = 0;
        subtotal_winlose_temp = 0;
        const res = await fetch(path_api+"api/companyinvoicemembertemp", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let css_totalwin = "";
                for (var i = 0; i < record.length; i++) {
                    if (parseInt(record[i]["totalwin"]) > 0) {
                        css_totalwin = "text-blue-700";
                    } else {
                        css_totalwin = "text-red-500";
                    }
                    subtotal_bet_temp = subtotal_bet_temp + parseInt(record[i]["totalbet"])
                    subtotal_bayar_temp = subtotal_bayar_temp + parseInt(record[i]["totalbayar"])
                    subtotal_cancel_temp = subtotal_cancel_temp + parseInt(record[i]["totalcancelbet"])
                    subtotal_win_temp = subtotal_win_temp + parseInt(record[i]["totalwin"])
                    if(subtotal_win_temp > 0){
                        subtotal_win_css_temp = "text-red-500";
                    }else{
                        subtotal_win_css_temp = "text-blue-700";
                    }
                    listInvoiceMembertemp = [
                        ...listInvoiceMembertemp,
                        {
                            member: record[i]["member"],
                            totalbet: record[i]["totalbet"],
                            totalbayar: record[i]["totalbayar"],
                            totalcancelbet: record[i]["totalcancelbet"],
                            totalwin: record[i]["totalwin"],
                            totalwin_css: css_totalwin,
                        },
                    ];
                }
                subtotal_winlose_temp = parseInt(subtotal_bayar_temp) - parseInt(subtotal_cancel_temp) - parseInt(subtotal_win_temp)
                if(subtotal_winlose_temp > 0){
                    subtotal_winlose_css_temp = "text-blue-700";
                }else{
                    subtotal_winlose_css_temp = "text-red-500";
                }
            }
        } else {
            // msgloader = json.message;
        }
        setTimeout(function () {
            // css_loader = "display: none;";
        }, 1000);
    }
    async function invoice_membersync() {
        dispatch("handleLoadingRunning", "call");
        const res = await fetch(path_api+"api/companyinvoicemembersync", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice_no),
            }),
        });
        const json = await res.json();
        if(!res.ok){
            let temp_msg = "System Mengalami Trouble"
            dispatch("handleLoadingRunningFinish", {
                    temp_msg
            });
        }else{
            let temp_msg = json.message
            dispatch("handleLoadingRunningFinish", {
                temp_msg
            });
            invoice_membertemp(invoice_no)
        }
    }
    async function invoice_listbet(e) {
        listInvoicelistbet = [];
        const res = await fetch(path_api+"api/companyinvoicelistpermainan", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice_no),
                permainan: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let bet_statuscss = ""
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["bet_status"] == "WINNER"){
                        bet_statuscss = "bg-[#ebfbee] text-[#6ec07b]"
                    }else if(record[i]["bet_status"] == "RUNNING"){
                        bet_statuscss = "bg-[#FFEB3B] text-black "
                    }else{
                        bet_statuscss = "bg-[#fde3e3] text-[#ea7779]"
                    }
                    listInvoicelistbet = [
                        ...listInvoicelistbet,
                        {
                            bet_id: record[i]["bet_id"].toString(),
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: bet_statuscss,
                        },
                    ];
                }
            }
        } 
    }
    async function invoice_grouppermainan(e) {
        listInvoiceGroupPermainan = []
        listInvoicelistbet = []
     
        const res = await fetch(path_api+"api/companyinvoicegrouppermainan", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice_no),
                username: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listInvoiceGroupPermainan = [
                        ...listInvoiceGroupPermainan,
                        {
                            permainan: record[i]["permainan"],
                        },
                    ];
                }
            }
        }
    }
    async function invoice_permainanstatus(e) {
        listInvoicelistbet = []
        const res = await fetch(path_api+"api/companyinvoicelistpermainanstatus", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice_no),
                status: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let bet_statuscss = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["bet_status"] == "WINNER"){
                        bet_statuscss = "bg-[#ebfbee] text-[#6ec07b]"
                    }else if(record[i]["bet_status"] == "RUNNING"){
                        bet_statuscss = "bg-[#FFEB3B] text-black "
                    }else{
                        bet_statuscss = "bg-[#fde3e3] text-[#ea7779]"
                    }
                    listInvoicelistbet = [
                        ...listInvoicelistbet,
                        {
                            bet_id: record[i]["bet_id"].toString(),
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: bet_statuscss,
                        },
                    ];
                }
            }
        }
    }
    const call_infoinvoice = (e) => {
        invoice_no = e;
        isModal_info_invoice = true;
        modal_infoinvoice_width = "max-w-7xl"
        invoice_member(e)
        invoice_membertemp(e)
    };
    const handleSelectPermainan = (event) => {
        invoice_listbet(event.target.value)
    };
    let tab_member = "bg-sky-600 text-white"
    let tab_allbet = ""
    let tab_winner = ""
    let tab_cancel = ""
    let panel_member = true
    let panel_allbet = false
    let panel_winner = false
    let panel_cancel = false
    const ChangeTabMenu_Invoice = (e) => {
        switch(e){
            case "menu_member":
                tab_member = "bg-sky-600 text-white"
                tab_allbet = ""
                tab_winner = ""
                tab_cancel = ""
                panel_member = true
                panel_allbet = false
                panel_winner = false
                panel_cancel = false
                break;
            case "menu_allbet":
                invoice_grouppermainan();
                tab_member = ""
                tab_allbet = "bg-sky-600 text-white"
                tab_winner = ""
                tab_cancel = ""
                panel_member = false
                panel_allbet = true
                panel_winner = false
                panel_cancel = false
                break;
            case "menu_listbetwinner":
                invoice_permainanstatus("WINNER");
                tab_member = ""
                tab_allbet = ""
                tab_winner = "bg-sky-600 text-white"
                tab_cancel = ""
                panel_member = false
                panel_allbet = false
                panel_winner = true
                panel_cancel = false
                break;
            case "menu_listbetcancel":
                invoice_permainanstatus("CANCEL");
                tab_member = ""
                tab_allbet = ""
                tab_winner = ""
                tab_cancel = "bg-sky-600 text-white"
                panel_member = false
                panel_allbet = false
                panel_winner = false
                panel_cancel = true
                break;
        }
    }
    $: {
        
        if(searchlistbet) {
            filterlistbet = listInvoicelistbet.filter(
                (item) =>
                    item.bet_status
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase()) ||
                    item.bet_nomortogel
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase()) ||
                    item.bet_username
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase()) || 
                    item.bet_id
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase())
            );
        } else {
            filterlistbet = [...listInvoicelistbet];
        }
    }
</script>
<div class="flex flex-col w-full">
    <div class="flex justify-center items-center w-full gap-2">
        <div class="w-full">
            <select
                bind:value="{select_periode}" 
                class="select select-bordered w-full rounded-sm focus:ring-0 focus:outline-none">
                <option disabled selected value="">--Pilih Periode--</option>
                <option value="01">JANUARY</option>
                <option value="02">FEBUARY</option>
                <option value="03">MARET</option>
                <option value="04">APRIL</option>
                <option value="05">MAY</option>
                <option value="06">JUNE</option>
                <option value="07">JULY</option>
                <option value="08">AUGUSTUS</option>
                <option value="09">SEPTEMBER</option>
                <option value="10">OCTOBER</option>
                <option value="11">NOVEMBER</option>
                <option value="12">DECEMBER</option>
            </select>
        </div>
        <div class="w-full">
            <select
                bind:value="{select_pasaran}" 
                class="select select-bordered w-full rounded-sm focus:ring-0 focus:outline-none">
                <option disabled selected value="">--Pilih Pasaran--</option>
                {#each listMasterPasaran as rec }
                    <option value="{rec.company_idcomppasaran}">{rec.company_nmpasarantogel}</option>
                {/each}
            </select>
        </div>
        <div class="w-full">
            <Button_custom 
                on:click={() => {
                    generate();
                }}
                button_style="btn-warning"
                button_disable={buttonLoading_flag}
                button_class=""
                button_disable_class="{buttonLoading_class}"
                button_title="Generate" />
        </div>
    </div>
    <div class="w-full scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[500px] overflow-y-scroll mt-2">
        <table class="table table-compact w-full ">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">NO</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">STATUS</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-center">DATE</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-left">INVOICE</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-left">PERIODE</th>
                    <th width="*" class="bg-[#475289] {font_size} text-white text-left">PASARAN</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-left">KELUARAN</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-right">REVISI</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-right">MEMBER</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-right">BET</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-right">BAYAR</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-right">CANCEL</th>
                    <th width="20%" class="bg-[#475289] {font_size} text-white text-right">WINLOSE</th>
                    <th width="20%" class="bg-[#475289] {font_size} text-white text-right">WINLOSE TEMP</th>
                    <th width="20%" class="bg-[#475289] {font_size} text-white text-right">SELISIH</th>
                </tr>
            </thead>
            {#if listPeriode != ""}
                    <tbody>
                        {#each listPeriode as rec}
                        <tr>
                            <td class="{font_size} align-top text-center">{rec.no}</td>
                            <td class="{font_size} align-top text-center">
                                <span class="{rec.company_pasaran_status_css} text-center rounded-md p-1 px-2 ">{rec.company_pasaran_status}</span>
                            </td>
                            <td class="{font_size} align-top text-center">{rec.company_pasaran_tanggal}</td>
                            <td on:click={() => {
                                call_infoinvoice(rec.company_pasaran_invoice);
                                }} class="{font_size} align-top text-left cursor-pointer underline">
                                {rec.company_pasaran_invoice}
                            </td>
                            <td class="{font_size} align-top text-left">{rec.company_pasaran_periode}</td>
                            <td class="{font_size} align-top text-left">{rec.company_pasaran_name}</td>
                            <td class="{font_size} align-top text-left">{rec.company_pasaran_keluaran}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_revisi_css}">{new Intl.NumberFormat().format(rec.company_pasaran_revisi)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_totalmember_css}">{new Intl.NumberFormat().format(rec.company_pasaran_totalmember)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_totalbet_css}">{new Intl.NumberFormat().format(rec.company_pasaran_totalbet)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_totaloutstanding_css}">{new Intl.NumberFormat().format(rec.company_pasaran_totaloutstanding)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_totalcancelbet_css}">{new Intl.NumberFormat().format(rec.company_pasaran_totalcancelbet)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_winlose_css}">{new Intl.NumberFormat().format(rec.company_pasaran_winlose)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_winlosetemp_css}">{new Intl.NumberFormat().format(rec.company_pasaran_winlosetemp)}</td>
                            <td class="{font_size} align-top text-right {rec.company_pasaran_selisih_css}">{new Intl.NumberFormat().format(rec.company_pasaran_selisih)}</td>
                            
                        </tr>
                        {/each}
                    </tbody>
                {:else}
                    <tbody>
                        <tr>
                            <td colspan="16" class="text-center">
                                <progress class="self-start progress progress-primary w-56"></progress>
                            </td>
                        </tr>
                    </tbody>
                {/if}
        </table>
    </div>
    <div class="bg-[#F7F7F7] rounded-sm h-10 p-2">
        <table class=" w-full">
            <tr>
                <td class="text-xs font-semibold text-left align-top">TOTAL WINLOSE</td>
                <td class="text-xs font-semibold text-right align-top {totalwinlose_css}">{new Intl.NumberFormat().format(totalwinlose)}</td>
            </tr>
        </table>
    </div>
</div>

<input type="checkbox" id="my-modal-infoinvoice" class="modal-toggle" bind:checked={isModal_info_invoice}>
<Modal_popup
    modal_popup_id="my-modal-infoinvoice"
    modal_popup_title="INVOICE : {invoice_no}"
    modal_popup_class="w-11/12 {modal_infoinvoice_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        <ul class="flex justify-center items-center gap-2">
            <li on:click={() => {
                    ChangeTabMenu_Invoice("menu_member");
                }}
                class="items-center {tab_member} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Member</li>
            <li on:click={() => {
                    ChangeTabMenu_Invoice("menu_allbet");
                }}
                class="items-center {tab_allbet} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">All Bet</li>
            <li on:click={() => {
                    ChangeTabMenu_Invoice("menu_listbetwinner");
                }}
                class="items-center {tab_winner} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Winner</li>
            <li on:click={() => {
                    ChangeTabMenu_Invoice("menu_listbetcancel");
                }}
                class="items-center {tab_cancel} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Cancel</li>
        </ul>
        {#if panel_member}
            {#if subtotal_winlose != subtotal_winlose_temp}
                <div class="flex justify-center items-center mt-2">
                    <Button_custom 
                        on:click={() => {
                            invoice_membersync();
                        }}
                        button_style="btn-warning"
                        button_disable={buttonLoading_flag}
                        button_class=""
                        button_disable_class="{buttonLoading_class}"
                        button_title="Sinkronisasi" />
                </div>
            {/if}
            <div class="grid grid-cols-2 gap-2 w-full">
                <div class="flex flex-col w-full gap-1">
                    <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[300px] overflow-y-scroll mt-2">
                        <table class="table table-compact w-full ">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="*" class="bg-[#475289] text-[11px] text-white text-left">MEMBER</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL BET</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL BAYAR</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL CANCEL</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL WIN</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listInvoiceMember as rec}
                                <tr>
                                    <td class="text-[11px] align-top text-left">{rec.member}</td>
                                    <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.totalbet)}</td>
                                    <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.totalbayar)}</td>
                                    <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.totalcancelbet)}</td>
                                    <td class="text-[11px] align-top text-right {rec.totalwin_css}">{new Intl.NumberFormat().format(rec.totalwin)}</td>
                                </tr>
                                {/each}
                            </tbody>
                        </table>
                    </div>
                    <div class="bg-[#F7F7F7] rounded-sm h-30 p-2">
                        <table class=" w-full">
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL BET</td>
                                <td class="text-[11px] font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_bet)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL BAYAR</td>
                                <td class="text-[11px] font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_bayar)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL CANCEL</td>
                                <td class="text-[11px] font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_cancel)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL WIN</td>
                                <td class="text-[11px] font-semibold text-right align-top {subtotal_win_css}">{new Intl.NumberFormat().format(subtotal_win)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL WINLOSE</td>
                                <td class="text-[11px] font-semibold text-right align-top {subtotal_winlose_css}">{new Intl.NumberFormat().format(subtotal_winlose)}</td>
                            </tr>
                        </table>
                    </div>
                </div>
                <div class="flex flex-col w-full gap-1">
                    <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[300px] overflow-y-scroll mt-2">
                        <table class="table table-compact w-full ">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="*" class="bg-[#475289] text-[11px] text-white text-left">MEMBER</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL BET</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL BAYAR</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL CANCEL</th>
                                    <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">TOTAL WIN</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listInvoiceMembertemp as rec}
                                <tr>
                                    <td class="text-[11px] align-top text-left">{rec.member}</td>
                                    <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.totalbet)}</td>
                                    <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.totalbayar)}</td>
                                    <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.totalcancelbet)}</td>
                                    <td class="text-[11px] align-top text-right {rec.totalwin_css}">{new Intl.NumberFormat().format(rec.totalwin)}</td>
                                </tr>
                                {/each}
                            </tbody>
                        </table>
                    </div>
                    <div class="bg-[#F7F7F7] rounded-sm h-30 p-2">
                        <table class=" w-full">
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL BET</td>
                                <td class="text-[11px] font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_bet_temp)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL BAYAR</td>
                                <td class="text-[11px] font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_bayar_temp)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL CANCEL</td>
                                <td class="text-[11px] font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_cancel_temp)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL WIN</td>
                                <td class="text-[11px] font-semibold text-right align-top {subtotal_win_css_temp}">{new Intl.NumberFormat().format(subtotal_win_temp)}</td>
                            </tr>
                            <tr>
                                <td class="text-[11px] font-semibold text-left align-top">SUBTOTAL WINLOSE</td>
                                <td class="text-[11px] font-semibold text-right align-top {subtotal_winlose_css_temp}">{new Intl.NumberFormat().format(subtotal_winlose_temp)}</td>
                            </tr>
                        </table>
                    </div>
                </div>
                
            </div>
        {/if}
        {#if panel_allbet}
            <div class="flex flex-col w-full gap-1 mt-2">
                <div class="flex gap-2">
                    <div class="w-full">
                        <select
                            on:change="{handleSelectPermainan}" 
                            class="select select-bordered w-full rounded-sm focus:ring-0 focus:outline-none">
                            <option disabled selected value="">--Pilih Periode--</option>
                            {#each listInvoiceGroupPermainan as rec }
                            <option value="{rec.permainan}">{rec.permainan}</option>
                            {/each}
                        </select>
                    </div>
                    <input 
                        bind:value={searchlistbet}
                        type="text" placeholder="Search" class="input input-bordered w-full  rounded-sm  focus:ring-0 focus:outline-none">
                </div>
                <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[500px] overflow-y-scroll mt-2">
                    <table class="table table-compact w-full ">
                        <thead class="sticky top-0">
                            <tr>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-center">STATUS</th>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-left">CODE</th>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-center">TANGGAL</th>
                                <th width="*" class="bg-[#475289] text-[11px] text-white text-left">USERNAME</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">IPADDRESS</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">DEVICE</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">PERMAINAN</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">NOMOR</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">BET</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">DISC</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">KEI</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">BAYAR</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">WIN</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">WIN TOTAL</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each filterlistbet as rec}
                            <tr>
                                <td class="text-[11px] align-top text-center">
                                    <span class="{rec.bet_statuscss} text-center rounded-md p-1 px-2 ">{rec.bet_status}</span>
                                </td>
                                <td class="text-[11px] align-top text-left">{rec.bet_id}</td>
                                <td class="text-[11px] align-top text-center">{rec.bet_datetime}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_username}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_ipaddress}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_device}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_typegame}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_nomortogel}</td>
                                <td class="text-[11px] align-top text-right">{rec.bet_bet}</td>
                                <td class="text-[11px] align-top text-right text-red-500">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_win)}</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
                            </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        {/if}
        {#if panel_winner}
            <div class="flex flex-col w-full gap-1 mt-2">
                <div class="flex gap-2">
                    <input 
                        bind:value={searchlistbet}
                        type="text" placeholder="Search" class="input input-bordered w-full  rounded-sm  focus:ring-0 focus:outline-none">
                </div>
                <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[500px] overflow-y-scroll mt-2">
                    <table class="table table-compact w-full ">
                        <thead class="sticky top-0">
                            <tr>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-center">STATUS</th>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-left">CODE</th>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-center">TANGGAL</th>
                                <th width="*" class="bg-[#475289] text-[11px] text-white text-left">USERNAME</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">IPADDRESS</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">DEVICE</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">PERMAINAN</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">NOMOR</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">BET</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">DISC</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">KEI</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">BAYAR</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">WIN</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">WIN TOTAL</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each filterlistbet as rec}
                            <tr>
                                <td class="text-[11px] align-top text-center">
                                    <span class="{rec.bet_statuscss} text-center rounded-md p-1 px-2 ">{rec.bet_status}</span>
                                </td>
                                <td class="text-[11px] align-top text-left">{rec.bet_id}</td>
                                <td class="text-[11px] align-top text-center">{rec.bet_datetime}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_username}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_ipaddress}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_device}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_typegame}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_nomortogel}</td>
                                <td class="text-[11px] align-top text-right">{rec.bet_bet}</td>
                                <td class="text-[11px] align-top text-right text-red-500">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                <td class="text-[11px] align-top text-right ">x{new Intl.NumberFormat().format(rec.bet_win)}</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
                            </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        {/if}
        {#if panel_cancel}
            <div class="flex flex-col w-full gap-1 mt-2">
                <div class="flex gap-2">
                    <input 
                        bind:value={searchlistbet}
                        type="text" placeholder="Search" class="input input-bordered w-full  rounded-sm  focus:ring-0 focus:outline-none">
                </div>
                <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[500px] overflow-y-scroll mt-2">
                    <table class="table table-compact w-full ">
                        <thead class="sticky top-0">
                            <tr>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-center">STATUS</th>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-left">CODE</th>
                                <th width="5%" class="bg-[#475289] text-[11px] text-white text-center">TANGGAL</th>
                                <th width="*" class="bg-[#475289] text-[11px] text-white text-left">USERNAME</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">IPADDRESS</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">DEVICE</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">PERMAINAN</th>
                                <th width="7%" class="bg-[#475289] text-[11px] text-white text-left">NOMOR</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">BET</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">DISC</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">KEI</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">BAYAR</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">WIN</th>
                                <th width="10%" class="bg-[#475289] text-[11px] text-white text-right">WIN TOTAL</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each filterlistbet as rec}
                            <tr>
                                <td class="text-[11px] align-top text-center">
                                    <span class="{rec.bet_statuscss} text-center rounded-md p-1 px-2 ">{rec.bet_status}</span>
                                </td>
                                <td class="text-[11px] align-top text-left">{rec.bet_id}</td>
                                <td class="text-[11px] align-top text-center">{rec.bet_datetime}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_username}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_ipaddress}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_device}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_typegame}</td>
                                <td class="text-[11px] align-top text-left">{rec.bet_nomortogel}</td>
                                <td class="text-[11px] align-top text-right">{rec.bet_bet}</td>
                                <td class="text-[11px] align-top text-right text-red-500">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                <td class="text-[11px] align-top text-right ">x{new Intl.NumberFormat().format(rec.bet_win)}</td>
                                <td class="text-[11px] align-top text-right text-blue-700">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
                            </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        {/if}
    </slot:template>
</Modal_popup>