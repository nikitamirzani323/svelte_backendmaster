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
    export let pasaran_minbet_kombinasi_field = 0;
    export let pasaran_maxbet_kombinasi_field = 0;
    export let pasaran_maxbuy_kombinasi_field = 0;
    export let pasaran_win_kombinasi_field = 0;
    export let pasaran_disc_kombinasi_field = 0;
    export let pasaran_limitglobal_kombinasi_field = 0;
    export let pasaran_limittotal_kombinasi_field = 0;
    let buttonLoading_flag = false;
    let buttonLoading_class = "btn-block";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_kombinasi_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_kombinasi_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_limittotal_kombinasi_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_kombinasi_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_win_kombinasi_field == "") {
            flag = true;
            msg_error += "The Win is required<br>";
        }
        if (pasaran_disc_kombinasi_field == "") {
            flag = true;
            msg_error += "The Diskon is required<br>";
        }
        if (flag == false) {
            buttonLoading_flag = true;
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/updatecompanypasarankombinasi", {
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
                    pasaran_minbet_kombinasi: parseInt(pasaran_minbet_kombinasi_field),
                    pasaran_maxbet_kombinasi: parseInt(pasaran_maxbet_kombinasi_field),
                    pasaran_maxbuy_kombinasi: parseInt(pasaran_maxbuy_kombinasi_field),
                    pasaran_limittotal_kombinasi: parseInt(pasaran_limittotal_kombinasi_field),
                    pasaran_limitglobal_kombinasi: parseInt(pasaran_limitglobal_kombinasi_field),
                    pasaran_win_kombinasi: parseFloat(pasaran_win_kombinasi_field),
                    pasaran_disc_kombinasi: parseFloat((pasaran_disc_kombinasi_field / 100).toPrecision(3)),
                }),
            });
            const json = await res.json();
            if(!res.ok){
                let temp_msg = "System Mengalami Trouble"
                dispatch("handleLoadingRunningFinish", {
                        temp_msg,idcompany,companypasaran_id
                });
            }else{
                let temp_msg = json.message
                dispatch("handleLoadingRunningFinish", {
                        temp_msg,idcompany,companypasaran_id
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
            const res = await fetch(path_api+"api/fetchpasaranmacau", {
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
                        temp_msg,idcompany,companypasaran_id
                });
            }else{
                let temp_msg = json.message
                dispatch("handleLoadingRunningFinish", {
                        temp_msg,idcompany,companypasaran_id
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
        bind:value={pasaran_minbet_kombinasi_field}
        input_id="pasaran_minbet_kombinasi_field"
        input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_limittotal_kombinasi_field}
        input_id="pasaran_limittotal_kombinasi_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_win_kombinasi_field}
        input_id="pasaran_win_kombinasi_field"
        input_placeholder="WIN(x)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_disc_kombinasi_field}
        input_id="pasaran_disc_kombinasi_field"
        input_placeholder="DISC(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_maxbet_kombinasi_field}
        input_id="pasaran_maxbet_kombinasi_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_limitglobal_kombinasi_field}
        input_id="pasaran_limitglobal_kombinasi_field"
        input_placeholder="Limit Global"/>
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_maxbuy_kombinasi_field}
        input_id="pasaran_maxbuy_kombinasi_field"
        input_placeholder="Max Buy"/>
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