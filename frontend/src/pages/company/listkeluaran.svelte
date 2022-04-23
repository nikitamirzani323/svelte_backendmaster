<script>
    import Button_custom from '../../components/button_custom.svelte'
    
    export let path_api = "";
    export let font_size = "";
    export let token = "";
    export let idcompany = "";
    export let companypasaran_id = "";
    export let listMasterPasaran = [];
    let listPeriode = [];
    let select_periode = "";
    let select_pasaran = "";
    let totalrecord = 0;
    let totalwinlose = 0;
    let totalwinlose_css = "";
    let buttonLoading_flag = false;
    let buttonLoading_class = "";
    let msg_error = "";
    async function generate() {
        listPeriode = [];
        totalwinlose = 0;
        totalwinlose_css = "text-blue-700";
        let flag = false;
        msg_error = "";
        
        if (select_periode == "") {
            flag = true;
            msg_error += "The Periode is required<br>";
        }
        if (select_pasaran == "") {
            flag = true;
            msg_error += "The Pasaran is required<br>";
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
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
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
</script>
<div class="flex flex-col w-full">
    <div class="flex justify-center items-center w-full gap-2">
        <div class="w-full">
            <select
                bind:value="{select_periode}" 
                class="select select-bordered w-full focus:ring-0 focus:outline-none">
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
                class="select select-bordered w-full focus:ring-0 focus:outline-none">
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
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center"></th>
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
                            <td  class="text-center text-xs cursor-pointer">
                               
                            </td>
                            <td class="{font_size} align-top text-center">{rec.no}</td>
                            <td class="{font_size} align-top text-center">
                                <span class="{rec.company_pasaran_status_css} text-center rounded-md p-1 px-2 ">{rec.company_pasaran_status}</span>
                            </td>
                            <td class="{font_size} align-top text-center">{rec.company_pasaran_tanggal}</td>
                            <td class="{font_size} align-top text-left">{rec.company_pasaran_invoice}</td>
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