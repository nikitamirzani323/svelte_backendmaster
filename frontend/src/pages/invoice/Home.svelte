<script>
    import { createEventDispatcher } from "svelte";
    import Button_custom from '../../components/button_custom.svelte'
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 

    export let path_api = "";
    export let font_size = "";
    export let token = "";
    export let master = "";
    export let listHome = [];
    export let totalrecord = 0;
    
    let page = "Invoice";
    let sData = "New";
    let isModal_Form_New = false
    let isModal_Show_Pasaran = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_flag = false;
    let buttonLoading_class = "btn-block";
    let msg_error = "";
    let select_periode = "";
    let searchHome = "";
    let filterHome = [];
    let idinvoice = "";
    let listinvoicedetail = [];
    let data_company = "";
    let data_periode = "";
    let total_winlose = 0;
    let total_winlose_class = "";
    let profitpasaran = 0;
    let profitpasaran_agen = 0;
    let totalprofitpasaran_all = 0;
    let totalprofitpasaran_agen_all = 0;
    let profitpasaran_class = "";
    let profitpasaran_agen_class = "";
    let totalprofitpasaran_all_class = "";
    let totalprofitpasaran_agen_all_class = "";
    let dispatch = createEventDispatcher();
   
    async function SaveTransaksi() {
        let flag = true;
        msg_error = "";
        if (select_periode == "") {
            flag = false
            msg_error += "The Periode is required<br>"
        }
        if (flag) {
            buttonLoading_flag = true;
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveinvoice", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    tipe: "",
                    master: master,
                    periode: select_periode,
                }),
            });
            const json = await res.json();
            if(!res.ok){
                loader_msg = "System Mengalami Trouble"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }else{
                if (json.status == 200) {
                    loader_msg = json.message
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_flag = false;
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true
            }
        }
    }
   
   
    async function UpdateWinlose(e) {
        buttonLoading_flag = true;
        loader_class = "inline-block"
        loader_msg = "Sending..."
        const res = await fetch(path_api+"api/saveinvoicewinlose", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                tipe: "UPDATE-WINLOSE",
                invoice: e,
                master: master,
            }),
        });
        const json = await res.json();
        if(!res.ok){
            loader_msg = "System Mengalami Trouble"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        }else{
            if (json.status == 200) {
                loader_msg = json.message
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            buttonLoading_flag = false;
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
            RefreshHalaman();
        }
    }
    async function UpdateStatus(e) {
        buttonLoading_flag = true;
        loader_class = "inline-block"
        loader_msg = "Sending..."
        const res = await fetch(path_api+"api/saveinvoicewinlose", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                tipe: "UPDATE-STATUS",
                invoice: e,
                master: master,
            }),
        });
        const json = await res.json();
        if(!res.ok){
            loader_msg = "System Mengalami Trouble"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        }else{
            if (json.status == 200) {
                loader_msg = json.message
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            buttonLoading_flag = false;
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
            RefreshHalaman();
        }
    }
    async function UpdateFetchPasaran() {
        buttonLoading_flag = true;
        loader_class = "inline-block"
        loader_msg = "Sending..."
        const res = await fetch(path_api+"api/saveinvoicewinlosepasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                tipe: "UPDATE-STATUS",
                invoice: idinvoice,
                master: master,
            }),
        });
        const json = await res.json();
        if(!res.ok){
            loader_msg = "System Mengalami Trouble"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        }else{
            if (json.status == 200) {
                loader_msg = json.message
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            buttonLoading_flag = false;
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
            call_invoicedetail(idinvoice);
            RefreshHalaman();
        }
    }
    async function DeleteFetchPasaran() {
        buttonLoading_flag = true;
        loader_class = "inline-block"
        loader_msg = "Sending..."
        const res = await fetch(path_api+"api/deleteinvoicewinlosepasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                tipe: "UPDATE-STATUS",
                invoice: idinvoice,
                master: master,
            }),
        });
        const json = await res.json();
        if(!res.ok){
            loader_msg = "System Mengalami Trouble"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        }else{
            if (json.status == 200) {
                loader_msg = json.message
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            buttonLoading_flag = false;
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
            call_invoicedetail(idinvoice);
            RefreshHalaman();
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        isModal_Form_New = true;
    };
    const ShowPasaran = (e,f,d) => {
        isModal_Show_Pasaran = true;
        idinvoice = e
        data_company = d
        data_periode = f
        call_invoicedetail(e)
        
    };
    async function call_invoicedetail(e) {
        listinvoicedetail = [];
        total_winlose = 0;
        const res = await fetch(path_api+"api/invoicedetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                invoice: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let winlose_class = "";
            profitpasaran = 0;
            profitpasaran_agen = 0;
            totalprofitpasaran_all = 0;
            totalprofitpasaran_agen_all = 0;
            profitpasaran_class = "";
            profitpasaran_agen_class = "";
            totalprofitpasaran_all_class = "";
            totalprofitpasaran_agen_all_class = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    total_winlose = total_winlose + record[i]["invoicedetail_winlose"]
                    if (record[i]["invoicedetail_winlose"] > 0) {
                        winlose_class = "text-blue-700"
                    } else {
                        winlose_class = "text-red-700"
                    }
                    if (total_winlose > 0) {
                        total_winlose_class = "text-blue-700"
                    } else {
                        total_winlose_class = "text-red-700"
                    }
                    let profitpasaran_temp = record[i]["invoicedetail_royaltyfee"] * record[i]["invoicedetail_winlose"]
                    
                    profitpasaran = profitpasaran_temp;
                    profitpasaran_agen = record[i]["invoicedetail_winlose"] - profitpasaran_temp;
                    totalprofitpasaran_all = totalprofitpasaran_all + profitpasaran
                    totalprofitpasaran_agen_all = totalprofitpasaran_agen_all + profitpasaran_agen

                    if (profitpasaran > 0) {
                        profitpasaran_class = "text-blue-700"
                    } else {
                        profitpasaran_class = "text-red-700"
                    }
                    if (profitpasaran_agen > 0) {
                        profitpasaran_agen_class = "text-blue-700"
                    } else {
                        profitpasaran_agen_class = "text-red-700"
                    }
                    listinvoicedetail = [
                        ...listinvoicedetail,
                        {
                            invoicedetail_id: record[i]["invoicedetail_id"],
                            invoicedetail_pasaran: record[i]["invoicedetail_pasaran"],
                            invoicedetail_winlose: record[i]["invoicedetail_winlose"],
                            invoicedetail_royaltyfee: (record[i]["invoicedetail_royaltyfee"]*100),
                            invoicedetail_profitcompany: profitpasaran,
                            invoicedetail_profitcompany_class: profitpasaran_class,
                            invoicedetail_profitagen: profitpasaran_agen,
                            invoicedetail_profitagen_class: profitpasaran_agen_class,
                            invoicedetail_create: record[i]["invoicedetail_create"],
                            invoicedetail_update: record[i]["invoicedetail_update"],
                            invoicedetail_winloseclass: winlose_class,
                        },
                    ];
                }
                totalprofitpasaran_agen_all_class = "text-red-600 font-semibold"
                totalprofitpasaran_all_class = "text-red-600 font-semibold"
                
                if(totalprofitpasaran_all > 0){
                    totalprofitpasaran_all_class = "text-blue-700 font-semibold"
                }
                if(totalprofitpasaran_agen_all > 0){
                    totalprofitpasaran_agen_all_class = "text-blue-700 font-semibold"
                }
            }
        }
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_company
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />

<Panel
    on:handleNewForm={NewData}
    on:handleRefresh={RefreshHalaman}
    panel_button_new={true}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_search">
        <div class="absolute inset-y-0 left-0 flex items-center pl-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
        </div>
        <input 
            bind:value={searchHome}
            type="text" placeholder="Search by Company, Periode" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center" colspan="3"></th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">NO</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">STATUS</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-center">CREATED</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-left">INVOICE</th>
                    <th width="*" class="bg-[#475289] {font_size} text-white text-left">COMPANY</th>
                    <th width="20%" class="bg-[#475289] {font_size} text-white text-left">PERIODE</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-right">TOTAL PASARAN</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-right">WINLOSE</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td 
                            title="Update Status"
                            class="text-center text-xs cursor-pointer">
                            {#if rec.home_status == "PROGRESS"}
                                <svg on:click={() => {
                                    UpdateStatus(rec.home_id);
                                    }} xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                    <path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                                </svg>
                            {/if}
                        </td>
                        <td 
                            title="Update winlose"
                            class="text-center text-xs cursor-pointer">
                            {#if rec.home_status == "PROGRESS"}
                                <svg 
                                    on:click={() => {
                                        UpdateWinlose(rec.home_id);
                                    }} xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                    <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                                </svg>
                            {/if}
                        </td>
                        <td 
                            title="Show List Pasaran"
                            class="text-center text-xs cursor-pointer">
                            {#if rec.home_status == "PROGRESS"}
                                <svg on:click={() => {
                                        ShowPasaran(rec.home_id,rec.home_name,rec.home_company);
                                    }}  xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                    <path stroke-linecap="round" stroke-linejoin="round" d="M9 13h6m-3-3v6m5 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                                </svg>
                            {/if}
                        </td>
                        <td class="{font_size} align-top text-center">{rec.home_no}</td>
                        <td class="{font_size} align-top text-center">
                            <span class="{rec.home_statusclass} text-center rounded-md p-1 px-2 ">{rec.home_status}</span>
                        </td>
                        <td class="{font_size} align-top text-center">{rec.home_create}</td>
                        <td class="{font_size} align-top text-left">{rec.home_id}</td>
                        <td class="{font_size} align-top text-left">{rec.home_company}</td>
                        <td class="{font_size} align-top text-left">{rec.home_name}</td>
                        <td class="{font_size} align-top text-right {rec.home_totalpasaranclass}">{new Intl.NumberFormat().format(rec.home_totalpasaran)}</td>
                        <td class="{font_size} align-top text-right {rec.home_winloseclass}">{new Intl.NumberFormat().format(rec.home_winlose)}</td>
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="10" class="text-center">
                            <progress class="self-start progress progress-primary w-56"></progress>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
    </slot:template>
</Panel>


<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title="Entry/{sData}"
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
            <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
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
            <div class="flex flex-wrap justify-end align-middle mt-2">
                <Button_custom 
                    on:click={() => {
                        SaveTransaksi();
                    }}
                    button_disable={buttonLoading_flag}
                    button_class="btn-block mt-2"
                    button_disable_class="{buttonLoading_class}"
                    button_title="Submit" />
            </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-showpasaran" class="modal-toggle" bind:checked={isModal_Show_Pasaran}>
<Modal_popup
    modal_popup_id="my-modal-showpasaran"
    modal_popup_title="{data_company} / {data_periode}"
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-col w-full mt-2">
            <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[450px] overflow-y-scroll mt-2">
                <table class="table table-compact w-full">
                    <thead class="sticky top-0">
                        <tr>
                            <th class="bg-[#475289] {font_size} text-white text-left">PASARAN</th>
                            <th class="bg-[#475289] {font_size} text-white text-right">WINLOSE</th>
                            <th class="bg-[#475289] {font_size} text-white text-right">FEE(%)</th>
                            <th class="bg-[#475289] {font_size} text-white text-right">COMPANY</th>
                            <th class="bg-[#475289] {font_size} text-white text-right">AGEN</th>
                        </tr>
                    </thead>
                    {#if listinvoicedetail != ""}
                    <tbody>
                        {#each listinvoicedetail as rec}
                        <tr>
                            <td class="{font_size} align-top text-left">{rec.invoicedetail_pasaran}</td>
                            <td class="{font_size} align-top text-right {rec.invoicedetail_winloseclass}">{new Intl.NumberFormat().format(rec.invoicedetail_winlose)}</td>
                            <td class="{font_size} align-top text-right {rec.invoicedetail_winloseclass}">{rec.invoicedetail_royaltyfee}</td>
                            <td class="{font_size} align-top text-right {rec.invoicedetail_profitcompany_class}">{new Intl.NumberFormat().format(rec.invoicedetail_profitcompany)}</td>
                            <td class="{font_size} align-top text-right {rec.invoicedetail_profitagen_class}">{new Intl.NumberFormat().format(rec.invoicedetail_profitagen)}</td>
                            
                        </tr>
                        {/each}
                    </tbody>
                {:else}
                    <tbody>
                        <tr>
                            <td colspan="10" class="text-center">
                                <progress class="self-start progress progress-primary w-56"></progress>
                            </td>
                        </tr>
                    </tbody>
                {/if}
                </table>
            </div>
            <div class="bg-[#eef4fb] rounded-sm h-16 p-2">
                <table class="w-full">
                    <tr>
                        <td class="text-xs text-left">TOTAL WINLOSE</td>
                        <td class="text-xs text-right {total_winlose_class}">{new Intl.NumberFormat().format(total_winlose)}</td>
                    </tr>
                    <tr>
                        <td class="text-xs text-left">TOTAL PROFIT COMPANY</td>
                        <td class="text-xs text-right {totalprofitpasaran_all_class}">{new Intl.NumberFormat().format(totalprofitpasaran_all)}</td>
                    </tr>
                    <tr>
                        <td class="text-xs text-left">TOTAL PROFIT AGEN</td>
                        <td class="text-xs text-right {totalprofitpasaran_agen_all_class}">{new Intl.NumberFormat().format(totalprofitpasaran_agen_all)}</td>
                    </tr>
                </table>
            </div>
            <div class="flex justify-center gap-2 w-full">
                <div class="w-full">
                    <Button_custom 
                        on:click={() => {
                            UpdateFetchPasaran();
                        }}
                        button_style="btn-warning"
                        button_disable={buttonLoading_flag}
                        button_class="btn-block mt-2"
                        button_disable_class="{buttonLoading_class}"
                        button_title="Fetch" />
                </div>
                <div class="w-full">
                    <Button_custom 
                        on:click={() => {
                            DeleteFetchPasaran();
                        }}
                        button_disable={buttonLoading_flag}
                        button_class="btn-block mt-2"
                        button_disable_class="{buttonLoading_class}"
                        button_title="Delete" />
                </div>
                
            </div>
        </div>
    </slot:template>
</Modal_popup>


<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



