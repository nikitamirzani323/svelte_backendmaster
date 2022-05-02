<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
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
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_flag = false;
    let buttonLoading_class = "btn-block";
    let msg_error = "";
    let searchHome = "";
    let filterHome = [];
    let idrecord = "";
    let domain_tipe_field = "";
    let domain_status_field = "";
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        domain_name_field: yup
            .string()
            .required("Domain is Required")
            .min(4, "Domain must be at least 4 Character")
            .max(250, "Domain must be at most 20 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            domain_name_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(values.domain_name_field);
        },
    });
    async function SaveTransaksi(domain) {
        let flag = true;
        msg_error = "";
        
        if(sData == "New"){
            if(domain_tipe_field == ""){
                flag = false
                msg_error += "The Tipe is required<br>"
            }
            if(domain_status_field == ""){
                flag = false
                msg_error += "The Status is required<br>"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg_error += "The ID is required<br>"
            }
            if(domain_tipe_field == ""){
                flag = false
                msg_error += "The Tipe is required<br>"
            }
            if(domain_status_field == ""){
                flag = false
                msg_error += "The Status is required<br>"
            }
        }
        if (flag) {
            buttonLoading_flag = true;
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savedomain", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master: master,
                    page:"DOMAIN-SAVE",
                    domain_id: parseInt(idrecord),
                    domain_name: $form.domain_name_field,
                    domain_tipe: domain_tipe_field,
                    domain_status: domain_status_field,
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
                    if(sData == "New"){
                        clearField();
                    }
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
   
    async function EditData(id,name,tipe,status) {
        msg_error = "";
        if(id != ""){
            sData = "Edit";
            clearField();
            
            idrecord = parseInt(id)
            $form.domain_name_field = name;
            domain_tipe_field = tipe;
            domain_status_field = status;
            isModal_Form_New = true;
        }else{
            isModalNotif = true;
            msg_error = "Silahkan Hubungi Administrator"
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        clearField()
        isModal_Form_New = true;
    };
    function clearField(){
        $form.domain_name_field = "";
        $errors.domain_name_field = "";
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_nama
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
            type="text" placeholder="Search by Domain" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center" colspan="2"></th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">NO</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">STATUS</th>
                    <th width="10%" class="bg-[#475289] {font_size} text-white text-center">CREATED</th>
                    <th width="*" class="bg-[#475289] {font_size} text-white text-left">COMPANY</th>
                    <th width="20%" class="bg-[#475289] {font_size} text-white text-left">PERIODE</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-right">WINLOSE</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td 
                            title="Update winlose"
                            on:click={() => {
                            EditData(rec.home_id,rec.home_nama,rec.home_tipe ,rec.home_status);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                              </svg>
                        </td>
                        <td 
                            title="Update Status"
                            on:click={() => {
                            EditData(rec.home_id,rec.home_nama,rec.home_tipe ,rec.home_status);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                            </svg>
                        </td>
                        <td class="{font_size} align-top text-center">{rec.home_no}</td>
                        <td class="{font_size} align-top text-center">
                            <span class="{rec.home_statusclass} text-center rounded-md p-1 px-2 ">{rec.home_status}</span>
                        </td>
                        <td class="{font_size} align-top text-center">{rec.home_create}</td>
                        <td class="{font_size} align-top text-left">{rec.home_company}</td>
                        <td class="{font_size} align-top text-left">{rec.home_name}</td>
                        <td class="{font_size} align-top text-right {rec.home_winloseclass}">
                            {new Intl.NumberFormat().format(rec.home_winlose)}
                        </td>
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
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_invalid={$errors.domain_name_field.length > 0}
                        bind:value={$form.domain_name_field}
                        input_id="domain_name_field"
                        input_placeholder="Domain"/>
                    {#if $errors.domain_name_field}
                        <small class="text-pink-600 text-[11px]">{$errors.domain_name_field}</small>
                    {/if}
                </div>
                <div class="relative form-control">
                    <select 
                        class="select select-bordered w-full rounded-sm" 
                        bind:value="{domain_tipe_field}">
                        <option disabled selected value="">--Pilih Tipe--</option>
                        <option value="FRONTEND">FRONTEND</option>
                        <option value="AGEN">AGEN</option>
                    </select>
                </div>
                <div class="relative form-control">
                    <select 
                        class="select select-bordered w-full rounded-sm" 
                        bind:value="{domain_status_field}">
                        <option disabled selected value="">--Pilih Status--</option>
                        <option value="RUNNING">RUNNING</option>
                        <option value="BANNED">BANNED</option>
                    </select>
                </div>
            </div>
            <div class="flex flex-wrap justify-end align-middle mt-2">
                <Button_custom 
                    on:click={() => {
                        handleSubmit();
                    }}
                    button_disable={buttonLoading_flag}
                    button_class="btn-block mt-2"
                    button_disable_class="{buttonLoading_class}"
                    button_title="Submit" />
            </div>
    </slot:template>
</Modal_popup>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



