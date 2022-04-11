<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 

    export let path_api = "";
    export let token = "";
    export let master = "";
    export let listHome = [];
    export let totalrecord = 0;

    let page = "Company";
    let sData = "New";
    let sDataAdmin = "New";
    let isModal_Form_New = false
    let isModal_Form_admin = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let modal_listadmin_width = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";
    let home_status_field = "";
    let home_create_field = "";
    let home_update_field = "";
    let idcompany = "";
    let listAdmin = [];
    let totalrecordadmin = 0;
    let tab_listadmin = "bg-sky-600 text-white"
    let tab_listpasaran = ""
    let panel_listadmin = true
    let panel_listpasaran = false
    let searchHome = "";
    let searchListAdmin = "";
    let filterHome = [];
    let filterListAdmin = [];

    let admin_username_field = "";
    let admin_username_enable_field = true;
    let admin_password_field = "";
    let admin_name_field = "";
    let admin_status_field = "";
    let admin_username_field_error = "";
    let admin_password_field_error = "";
    let admin_name_field_error = "";
    let admin_status_field_error = "";

    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        home_id_field: yup
            .string()
            .required("ID is Required")
            .matches(
                /^[A-z]+$/,
                "ID must Character A-Z "
            )
            .min(3, "ID must be at least 4 Character")
            .max(10, "ID must be at most 6 Character"),
        home_name_field: yup
            .string()
            .required("Name is Required")
            .matches(
                /^[A-z0-9 ]+$/,
                "Name must Character A-Z  or 1-9"
            )
            .min(4, "Name must be at least 4 Character")
            .max(70, "Name must be at most 70 Character"),
        home_url_field: yup
            .string()
            .required("Url is Required")
            .min(4, "Url must be at least 4 Character")
            .max(350, "Url must be at most 350 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            home_id_field: "",
            home_name_field: "",
            home_url_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(values.home_id_field,values.home_name_field,values.home_url_field);
        },
    });
    
    async function SaveTransaksi(idcomp,namecomp,urlcomp) {
        let flag = true;
        msg_error = "";
        
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompany", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master: master,
                    company: idcomp,
                    name: namecomp,
                    urldomain: urlcomp,
                    status: "ACTIVE",
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
                        $form.home_name_field = "";
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            alert(msg_error);
        }
    }
    async function EditData(e) {
        if(e != ""){
            isModalLoading = true;
            modal_width = "max-w-5xl";
            sData = "Edit";
            clearField();
            idcompany = e;
            $form.home_id_field = e;
            const res = await fetch(path_api+"api/companydetail", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    company: e,
                    page: "COMPANY_SAVE",
                    master: master,
                    sData: "Edit",
                }),
            });
            const json = await res.json();
            let record = json.record;
            if (json.status === 400) {
                isModalNotif = true;
                msg_error = "Maaf Saat Ini Anda Tidak Bisa Mengakses Halaman Ini"
            }else if(json.status === 200) {
                isModal_Form_New = true;
                isModalLoading = false;
                for (let i = 0; i < record.length; i++) {
                    $form.home_name_field = record[i]["company_name"];
                    $form.home_url_field = record[i]["company_url"];
                    home_status_field = record[i]["company_status"];
                    home_create_field = record[i]["company_create"];
                    home_update_field = record[i]["company_update"];
                }
                call_listadmin();
            }else{
                isModalLoading = false;
                isModalNotif = true;
                msg_error = "Silahkan Hubungi Administrator"
            }
        }
    }
    async function call_listadmin() {
        listAdmin = [];
        const res = await fetch(path_api+"api/companylistadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                sData: "New",
                page: "COMPANY_HOME",
                company: idcompany,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                totalrecordadmin = record.length;
                let company_admin_status_class = "";
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["company_admin_status"] == "ACTIVE"){
                        company_admin_status_class = "bg-[#8BC34A] text-black"
                    }else{
                        company_admin_status_class = "bg-red-600 text-white"
                    }
                    listAdmin = [
                        ...listAdmin,
                        {
                            company_admin_username:record[i]["company_admin_username"],
                            company_admin_typeadmin:record[i]["company_admin_typeadmin"],
                            company_admin_nama: record[i]["company_admin_nama"],
                            company_admin_status:record[i]["company_admin_status"],
                            company_admin_status_class:company_admin_status_class,
                            company_admin_lastlogin:record[i]["company_admin_lastlogin"],
                            company_admin_lastipaddres:record[i]["company_admin_lastipaddres"],
                            company_admin_create:record[i]["company_admin_create"],
                            company_admin_update:record[i]["company_admin_update"],
                        },
                    ];
                }
            }
        } 
    }
    async function handleSaveAdmin() {
        let flag = false;
        msg_error = "";
        admin_username_field_error = "";
        admin_password_field_error = "";
        admin_name_field_error = "";
        admin_status_field_error = "";
        const regexExp = /^[a-z0-9]+$/gi;
        const regexExp2 = /^[A-Za-z0-9 ]+$/gi;
        let flag_username_pattern = regexExp.test(admin_username_field)
        let flag_name_pattern = regexExp2.test(admin_name_field)
        if(admin_username_field == ""){
            flag = true
            admin_username_field_error ="Username is required"
        }
        if(!flag_username_pattern){
            flag = true
            admin_username_field_error ="Format Username tidak sesuai pattern a-z atau 0-9"
        }
        if(sDataAdmin == "New"){
            if(admin_password_field == ""){
                flag = true
                admin_password_field_error ="Password is required"
            }
        }
        if(admin_name_field == ""){
            flag = true
            admin_name_field_error ="Name is required"
        }
        if(!flag_name_pattern){
            flag = true
            admin_name_field_error ="Format Name tidak sesuai pattern a-z atau A-Z atau 0-9 atau spasi"
        }
        if(admin_status_field == ""){
            flag = true
            admin_status_field_error ="Status is required"
        }
        if(flag == false){
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompanyadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataAdmin,
                    master: master,
                    company: idcompany,
                    admin_username: admin_username_field.toLowerCase,
                    admin_password: admin_password_field,
                    admin_name: admin_name_field,
                    admin_status: admin_status_field,
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
                    if(sDataAdmin == "New"){
                        admin_username_field = "";
                        admin_password_field = "";
                        admin_name_field = "";
                        admin_status_field = "";
                        admin_username_field_error = "";
                        admin_password_field_error = "";
                        admin_name_field_error = "";
                        admin_status_field_error = "";
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                call_listadmin();
            }
        }
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        clearField()
        modal_width = "max-w-xl"
        isModal_Form_New = true;
    };
    const NewDataAdmin = (e,usernameadmincomp,nmadmincomp,stausadmincomp) => {
        sDataAdmin = e;
        modal_listadmin_width = "max-w-xl"
        isModal_Form_admin = true;
        if(e == "Edit"){
            admin_username_enable_field = false
            admin_username_field = usernameadmincomp;
            admin_password_field = "";
            admin_name_field = nmadmincomp;
            admin_status_field = stausadmincomp;
        }else{
            admin_username_enable_field = true
            admin_username_field = "";
            admin_password_field = "";
            admin_name_field = "";
            admin_status_field = "";
            admin_username_field_error = "";
            admin_password_field_error = "";
            admin_name_field_error = "";
            admin_status_field_error = "";
        }
    };
    
    const ChangeTabMenu = (e) => {
        switch(e){
            case "menu_listadmin":
                tab_listadmin = "bg-sky-600 text-white"
                tab_listpasaran = ""
                panel_listadmin = true
                panel_listpasaran = false
                break;
            case "menu_listpasaran":
                tab_listadmin = ""
                tab_listpasaran = "bg-sky-600 text-white"
                panel_listadmin = false
                panel_listpasaran = true
                break;
        }
    }
   
    function clearField(){
        idcompany = "";
        $form.home_id_field = "";
        $form.home_name_field = "";
        $form.home_url_field = "";
        $errors.home_id_field = "";
        $errors.home_name_field = "";
        $errors.home_url_field = "";
        home_status_field = "";
        home_create_field = "";
        home_update_field = "";
        tab_listadmin = "bg-sky-600 text-white"
        tab_listpasaran = ""
        panel_listadmin = true
        panel_listpasaran = false
        listAdmin = [];
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

        if (searchListAdmin) {
            filterListAdmin = listAdmin.filter(
                (item) =>
                    item.company_admin_username
                        .toLowerCase()
                        .includes(searchListAdmin.toLowerCase())
            );
        } else {
            filterListAdmin = [...listAdmin];
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
            type="text" placeholder="Search by Rule" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full ">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">&nbsp;</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">START</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">END</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">ID</th>
                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">COMPANY</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PERIODE</th>
                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WINLOSE</th>
                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WINLOSE_TEMP</th>
                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">SELISIH</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EditData(rec.home_idcompany);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_no}</td>
                        <td class="text-xs lg:text-sm align-top text-left">
                            <span class="{rec.home_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.home_status}</span>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_startjoin}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_endjoin}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_idcompany}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_name}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_periode}</td>
                        <td class="text-xs lg:text-sm align-top text-right {rec.home_winlose_class}">{new Intl.NumberFormat().format(rec.home_winlose)}</td>
                        <td class="text-xs lg:text-sm align-top text-right {rec.home_winlosetemp_class}">{new Intl.NumberFormat().format(rec.home_winlosetemp)}</td>
                        <td class="text-xs lg:text-sm align-top text-right {rec.home_selisihwinlose_class}">{new Intl.NumberFormat().format(rec.home_selisihwinlose)}</td>
                        
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="12" class="text-center">
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
    modal_popup_class="select-none w-11/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        {#if sData=="New"}
            <div class="flex flex-auto flex-col overflow-auto gap-3 mt-2 ">
                <div class="mt-2">
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={true}
                        input_required={true}
                        input_tipe="text"
                        input_text_class="uppercase"
                        input_maxlength_text="10"
                        input_invalid={$errors.home_id_field.length > 0}
                        bind:value={$form.home_id_field}
                        input_id="home_id_field"
                        input_placeholder="ID"/>
                    {#if $errors.home_id_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_id_field}</small>
                    {/if}
                </div>
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_text_class="uppercase"
                    input_maxlength_text="70"
                    input_invalid={$errors.home_name_field.length > 0}
                    bind:value={$form.home_name_field}
                    input_id="home_name_field"
                    input_placeholder="Name"/>
                {#if $errors.home_name_field}
                    <small class="text-pink-600 text-[11px]">{$errors.home_name_field}</small>
                {/if}
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="350"
                    input_invalid={$errors.home_url_field.length > 0}
                    bind:value={$form.home_url_field}
                    input_id="home_url_field"
                    input_placeholder="URL"/>
                {#if $errors.home_url_field}
                    <small class="text-pink-600 text-[11px]">{$errors.home_url_field}</small>
                {/if}
            </div>
            <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
                <button
                    on:click={() => {
                        handleSubmit();
                    }}  
                    class="{buttonLoading_class}">Submit</button>
            </div>
        {/if}
        {#if sData=="Edit"}
            <div class="flex justify-between w-full gap-2">
                <div class="w-1/2">
                    <div class="flex flex-auto flex-col overflow-auto gap-2 mt-2  ">
                        <div class="mt-2">
                            <Input_custom
                                input_autofocus={false}
                                input_required={true}
                                input_enabled={false}
                                input_tipe="text"
                                input_text_class="uppercase"
                                input_maxlength_text="10"
                                bind:value={$form.home_id_field}
                                input_id="home_id_field"
                                input_placeholder="ID"/>
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_text_class="uppercase"
                                input_maxlength_text="70"
                                input_invalid={$errors.home_name_field.length > 0}
                                bind:value={$form.home_name_field}
                                input_id="home_name_field"
                                input_placeholder="Name"/>
                            {#if $errors.home_name_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_name_field}</small>
                            {/if}
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_maxlength_text="350"
                                input_invalid={$errors.home_url_field.length > 0}
                                bind:value={$form.home_url_field}
                                input_id="home_url_field"
                                input_placeholder="URL"/>
                            {#if $errors.home_url_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_url_field}</small>
                            {/if}
                        </div>
                        <p class="text-[11px]">
                            Create : {home_create_field}
                            {#if home_update_field != ""}
                                <br>
                                Update : {home_update_field}
                            {/if}
                        </p>
                    </div>
                    <div class="flex flex-wrap justify-end align-middle  mt-2">
                        <button
                            on:click={() => {
                                handleSubmit();
                            }}  
                            class="{buttonLoading_class} btn-block">Submit</button>
                    </div>
                </div>
                <div class="w-full p-2">
                    <ul class="flex justify-center items-center gap-2">
                        <li on:click={() => {
                                ChangeTabMenu("menu_listadmin");
                            }}
                            class="items-center {tab_listadmin} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Admin</li>
                        <li on:click={() => {
                                ChangeTabMenu("menu_listpasaran");
                            }}
                            class="items-center {tab_listpasaran} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Pasaran</li>
                    </ul>
                    {#if panel_listadmin}
                        <h2 class="text-lg font-bold">List Admin</h2>
                        <div class="flex justify-between items-center w-full gap-1">
                            <input 
                                bind:value={searchListAdmin}
                                type="text" placeholder="Search by Username" class="input input-bordered w-full  rounded-md  focus:ring-0 focus:outline-none">
                            <button on:click={() => {
                                NewDataAdmin("New","","","");
                                }}  class="btn bg-primary hover:bg-primary  rounded-md shadow-lg shadow-[#6eb5d8] border-none  ">New</button>
                        </div>
                        <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[400px] overflow-y-scroll mt-2">
                            <table class="table table-compact w-full">
                                <thead class="sticky top-0">
                                    <tr>
                                        <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">&nbsp;</th>
                                        <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">STATUS</th>
                                        <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TYPE</th>
                                        <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">LASTLOGIN</th>
                                        <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">LASTIPADDRESS</th>
                                        <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#if filterListAdmin != ""}
                                        {#each filterListAdmin as rec}
                                            <tr>
                                                <td class="cursor-pointer" on:click={() => {
                                                        NewDataAdmin("Edit",rec.company_admin_username,rec.company_admin_nama,rec.company_admin_status);
                                                    }}>
                                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                                                    </svg>
                                                </td>
                                                <td class="text-xs text-center align-top">
                                                    <span class="{rec.company_admin_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.company_admin_status}</span>
                                                </td>
                                                <td class="text-xs text-left align-top">{rec.company_admin_typeadmin}</td>
                                                <td class="text-xs text-center align-top">{rec.company_admin_lastlogin}</td>
                                                <td class="text-xs text-center align-top">{rec.company_admin_lastipaddres}</td>
                                                <td class="text-xs text-left align-top">{rec.company_admin_username}</td>
                                            </tr>
                                        {/each}
                                    {:else}
                                        <tr>
                                            <td colspan="5" class="text-xs">No Records</td>
                                        </tr>
                                    {/if}
                                </tbody>
                            </table>
                        </div>
                        <div class="bg-[#F7F7F7] rounded-sm h-10 p-2">
                            <table class=" w-full">
                                <tr>
                                    <td class="text-xs font-semibold text-left align-top">TOTAL RECORD</td>
                                    <td class="text-xs font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(totalrecordadmin)}</td>
                                </tr>
                            </table>
                        </div>
                    {/if}
                </div>
            </div>
        {/if}
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formpasaran" class="modal-toggle" bind:checked={isModal_Form_admin}>
<Modal_popup
    modal_popup_id="my-modal-formpasaran"
    modal_popup_title="Admin Entry/{sDataAdmin}"
    modal_popup_class="select-none w-11/12 {modal_listadmin_width} scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-col gap-2">
            <div class="mt-2">
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_enabled={admin_username_enable_field}
                    input_tipe="text"
                    input_text_class="lowercase"
                    input_maxlength_text="30"
                    bind:value={admin_username_field}
                    input_id="admin_username_field"
                    input_placeholder="Username"/>
                {#if admin_username_field_error != ""}
                    <small class="text-pink-600 text-[11px]">{admin_username_field_error}</small>
                {/if}
            </div>
            <div>
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="password"
                    input_attr="password"
                    bind:value={admin_password_field}
                    input_id="admin_password_field"
                    input_placeholder="Password"/>
                    {#if admin_password_field_error != ""}
                        <small class="text-pink-600 text-[11px]">{admin_password_field_error}</small>
                    {/if}
            </div>
            <div>
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_enabled={true}
                    input_tipe="text"
                    input_text_class=""
                    input_maxlength_text="70"
                    bind:value={admin_name_field}
                    input_id="admin_name_field"
                    input_placeholder="Name"/>
                    {#if admin_name_field_error != ""}
                        <small class="text-pink-600 text-[11px]">{admin_name_field_error}</small>
                    {/if}
            </div>
            <div>
                <select class="select select-bordered w-full rounded-md focus:ring-0 focus:outline-none" bind:value="{admin_status_field}">
                    <option disabled selected value="">--Pilih Status--</option>
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="DEACTIVE">DEACTIVE</option>
                </select>
                {#if admin_status_field_error != ""}
                    <small class="text-pink-600 text-[11px]">{admin_status_field_error}</small>
                {/if}
            </div>
            <button
                on:click={() => {
                    handleSaveAdmin();
                }}  
                class="{buttonLoading_class} btn-block">Submit</button>
        </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />
