<script>
    import { createEventDispatcher } from "svelte";
    import Button_custom from '../../components/button_custom.svelte'
    import Input_custom from '../../components/Input.svelte'
    
    export let path_api = "";
    export let master = "";
    export let token = "";
    export let idcompany = "";
    export let companypasaran_id = "";
    export let pasaran_id_field = "";
    export let pasaran_minbet_cjitu_field = 0;
    export let pasaran_maxbet_cjitu_field = 0;
    export let pasaran_maxbuy_cjitu_field = 0;
    export let pasaran_winas_cjitu_field = 0;
    export let pasaran_winkop_cjitu_field = 0;
    export let pasaran_winkepala_cjitu_field = 0;
    export let pasaran_winekor_cjitu_field = 0;
    export let pasaran_desc_cjitu_field = 0;
    export let pasaran_limitglobal_cjitu_field = 0;
    export let pasaran_limittotal_cjitu_field = 0;
    let buttonLoading_flag = false;
    let buttonLoading_class = "btn-block";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_cjitu_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_cjitu_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_maxbuy_cjitu_field == "") {
            flag = true;
            msg_error += "The Max Buy is required<br>";
        }
        if (pasaran_limittotal_cjitu_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_cjitu_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_winas_cjitu_field == "") {
            flag = true;
            msg_error += "The Win AS is required<br>";
        }
        if (pasaran_winkop_cjitu_field == "") {
            flag = true;
            msg_error += "The Win KOP is required<br>";
        }
        if (pasaran_winkepala_cjitu_field == "") {
            flag = true;
            msg_error += "The Win KEPALA is required<br>";
        }
        if (pasaran_winekor_cjitu_field == "") {
            flag = true;
            msg_error += "The Win EKOR is required<br>";
        }
        if (pasaran_desc_cjitu_field == "") {
            flag = true;
            msg_error += "The Diskon is required<br>";
        }
        if (flag == false) {
            buttonLoading_flag = true;
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/updatecompanypasarancolokjitu", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "Edit",
                    master: master,
                    company: idcompany,
                    companypasaran_id: companypasaran_id,
                    pasaran_id: pasaran_id_field,
                    pasaran_minbet_cjitu: parseInt(pasaran_minbet_cjitu_field),
                    pasaran_maxbet_cjitu: parseInt(pasaran_maxbet_cjitu_field),
                    pasaran_maxbuy_cjitu: parseInt(pasaran_maxbuy_cjitu_field),
                    pasaran_limittotal_cjitu: parseInt(pasaran_limittotal_cjitu_field),
                    pasaran_limitglobal_cjitu: parseInt(pasaran_limitglobal_cjitu_field),
                    pasaran_winas_cjitu: parseFloat(pasaran_winas_cjitu_field),
                    pasaran_winkop_cjitu: parseFloat(pasaran_winkop_cjitu_field),
                    pasaran_winkepala_cjitu: parseFloat(pasaran_winkepala_cjitu_field),
                    pasaran_winekor_cjitu: parseFloat(pasaran_winekor_cjitu_field / 100),
                    pasaran_desc_cjitu: parseFloat(pasaran_desc_cjitu_field / 100),
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
            }
            buttonLoading_flag = false;
        } else {
            if(msg_error != ""){
                let temp_msg = msg_error
                dispatch("handleCallNotif", {
                        temp_msg
                });
            }
        }
    }
    async function fetchcolok() {
        let flag = false;
        msg_error = "";
        if (pasaran_id_field == "") {
            flag = true;
            msg_error += "The Pasaran is required<br>";
        }
        if (flag == false) {
            buttonLoading_flag = true;
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/fetchpasarancjitu", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "Edit",
                    master: master,
                    company: idcompany,
                    pasaran_id: pasaran_id_field,
                    Companypasaran_id: companypasaran_id,
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
            }
            buttonLoading_flag = false;
        } else {
            if(msg_error != ""){
                let temp_msg = msg_error
                dispatch("handleCallNotif", {
                        temp_msg
                });
            }
        }
    }
</script>

<div class="grid grid-cols-4 gap-1 mt-2 mb-5">
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_minbet_cjitu_field}
        input_id="pasaran_minbet_cjitu_field"
        input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_limittotal_cjitu_field}
        input_id="pasaran_limittotal_cjitu_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_winas_cjitu_field}
        input_precision=2
        input_id="pasaran_winas_cjitu_field"
        input_placeholder="WIN AS(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_desc_cjitu_field}
        input_precision=2
        input_id="pasaran_desc_cjitu_field"
        input_placeholder="DISC(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_maxbet_cjitu_field}
        input_id="pasaran_maxbet_cjitu_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_limitglobal_cjitu_field}
        input_id="pasaran_limitglobal_cjitu_field"
        input_placeholder="Limit Global"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_winkop_cjitu_field}
        input_precision=2
        input_id="pasaran_winkop_cjitu_field"
        input_placeholder="WIN KOP(%)"/>
    <div></div>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_maxbuy_cjitu_field}
        input_id="pasaran_maxbuy_cjitu_field"
        input_placeholder="Max Buy"/>
    <div ></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_winkepala_cjitu_field}
        input_precision=2
        input_id="pasaran_winkepala_cjitu_field"
        input_placeholder="WIN KEPALA(%)"/>
    <div></div>
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_winekor_cjitu_field}
        input_precision=2
        input_id="pasaran_winekor_cjitu_field"
        input_placeholder="WIN EKOR(%)"/>
</div>
<div class="grid grid-cols-2 gap-2">
    <Button_custom 
        on:click={() => {
            fetchcolok();
        }}
        button_style="btn-warning"
        button_disable={buttonLoading_flag}
        button_class="btn-block mt-2"
        button_disable_class="{buttonLoading_class}"
        button_title="Fetch" />
    <Button_custom 
        on:click={() => {
            save432d();
        }}
        button_disable={buttonLoading_flag}
        button_class="btn-block mt-2"
        button_disable_class="{buttonLoading_class}"
        button_title="Submit" />
</div>