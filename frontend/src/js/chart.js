import {ChartManager} from './chart_manager'
import {Control} from './control'
import Kline from './kline'
import {Template} from './templates'
import $ from 'jquery'

export class Chart {
    constructor() {
        this._data = null;
        this._charStyle = "CandleStick";
        this._depthData = {
            array: null,
            asks_count: 0,
            bids_count: 0,
            asks_si: 0,
            asks_ei: 0,
            bids_si: 0,
            bids_ei: 0
        };
        this.strIsLine = false;
        this._range = Kline.instance.range;
        this._symbol = Kline.instance.symbol;
    }

    setTitle() {
        let lang = ChartManager.instance.getLanguage();
        let title = Kline.instance.symbolName;
        title += ' ';
        title += this.strIsLine ? Chart.strPeriod[lang]['line'] : Chart.strPeriod[lang][this._range];
        title += (this._contract_unit + '/' + this._money_type).toUpperCase();
        ChartManager.instance.setTitle('frame0.k0', title);
    }


    setSymbol(symbol) {
        this._symbol = symbol;
        this.updateDataAndDisplay();
    }

    updateDataAndDisplay() {
        Kline.instance.symbol = this._symbol;
        Kline.instance.range = this._range;
        ChartManager.instance.setCurrentDataSource('frame0.k0', this._symbol + '.' + this._range);
        ChartManager.instance.setNormalMode();
        let f = Kline.instance.chartMgr.getDataSource("frame0.k0").getLastDate();
        if (f === -1) {
            Kline.instance.requestParam = Control.setHttpRequestParam(Kline.instance.symbol, Kline.instance.range, Kline.instance.limit, null);
            Control.requestData(true);
        } else {
            Kline.instance.requestParam = Control.setHttpRequestParam(Kline.instance.symbol, Kline.instance.range, null, f.toString());
            Control.requestData();
        }
        ChartManager.instance.redraw('All', false);
    }


    setCurrentContractUnit(contractUnit) {
        this._contract_unit = contractUnit;
        this.updateDataAndDisplay();
    }

    setCurrentMoneyType(moneyType) {
        this._money_type = moneyType;
        this.updateDataAndDisplay();
    }

    setCurrentPeriod(period) {
        this._range = Kline.instance.periodMap[period];
        if (Kline.instance.type === "stomp" && Kline.instance.stompClient.ws.readyState === 1) {
            Kline.instance.subscribed.unsubscribe();
            Kline.instance.subscribed = Kline.instance.stompClient.subscribe(Kline.instance.subscribePath + '/' + Kline.instance.symbol + '/' + this._range, Control.subscribeCallback);
        }
        this.updateDataAndDisplay();
        Kline.instance.onRangeChangeFunc(this._range);
    }

    updateDataSource(data) {
        this._data = data;
        ChartManager.instance.updateData("frame0.k0", this._data);
    }

    updateDepth(array) {
        if (array == null) {
            this._depthData.array = [];
            ChartManager.instance.redraw('All', false);
            return;
        }
        if (!array.asks || !array.bids || array.asks === '' || array.bids === '')
            return;
        let _data = this._depthData;
        _data.array = [];
        for (let i = 0; i < array.asks.length; i++) {
            let data = {};
            data.rate = array.asks[i][0];
            data.amount = array.asks[i][1];
            _data.array.push(data);
        }
        for (let i = 0; i < array.bids.length; i++) {
            let data = {};
            data.rate = array.bids[i][0];
            data.amount = array.bids[i][1];
            _data.array.push(data);
        }
        _data.asks_count = array.asks.length;
        _data.bids_count = array.bids.length;
        _data.asks_si = _data.asks_count - 1;
        _data.asks_ei = 0;
        _data.bids_si = _data.asks_count - 1;
        _data.bids_ei = _data.asks_count + _data.bids_count - 2;
        for (let i = _data.asks_si; i >= _data.asks_ei; i--) {
            if (i === _data.asks_si && _data.array[i] !== undefined) {
                _data.array[i].amounts = _data.array[i].amount;
            } else if(_data.array[i + 1] !== undefined) {
                _data.array[i].amounts = _data.array[i + 1].amounts + _data.array[i].amount;
            }
        }
        for (let i = _data.bids_si; i <= _data.bids_ei; i++) {
            if (i === _data.bids_si && _data.array[i] !== undefined) {
                _data.array[i].amounts = _data.array[i].amount;
            } else if (_data.array[i - 1] !== undefined) {
                _data.array[i].amounts = _data.array[i - 1].amounts + _data.array[i].amount;
            }
        }
        ChartManager.instance.redraw('All', false);
    }

    setMainIndicator(indicName) {
        this._mainIndicator = indicName;
        if (indicName === 'NONE') {
            ChartManager.instance.removeMainIndicator('frame0.k0');
        } else {
            ChartManager.instance.setMainIndicator('frame0.k0', indicName);
        }
        ChartManager.instance.redraw('All', true);
    }

    setIndicator(index, indicName) {
        if (indicName === 'NONE') {
            /*
            let index = 2;
            if (Template.displayVolume === false)
                index = 1;
            */
            let index=1;
            let areaName = ChartManager.instance.getIndicatorAreaName('frame0.k0', index);
            if (areaName !== '')
                ChartManager.instance.removeIndicator(areaName);
        } else {
            /*
            let index = 2;
            if (Template.displayVolume === false)
                index = 1;
            */
            let index=1;
            let areaName = ChartManager.instance.getIndicatorAreaName('frame0.k0', index);
            if (areaName === '') {
                Template.createIndicatorChartComps('frame0.k0', indicName);
            } else {
                ChartManager.instance.setIndicator(areaName, indicName);
            }
        }
        ChartManager.instance.redraw('All', true);
    }

    addIndicator(indicName) {
        ChartManager.instance.addIndicator(indicName);
        ChartManager.instance.redraw('All', true);
    }

    removeIndicator(indicName) {
        let areaName = ChartManager.instance.getIndicatorAreaName(2);
        ChartManager.instance.removeIndicator(areaName);
        ChartManager.instance.redraw('All', true);
    };

}
Chart.strPeriod = {
    'zh-cn': {
        'line': '(??????)',
        '1min': '(1??????)',
        '5min': '(5??????)',
        '15min': '(15??????)',
        '30min': '(30??????)',
        '1hour': '(1??????)',
        '1day': '(??????)',
        '1week': '(??????)',
        '3min': '(3??????)',
        '2hour': '(2??????)',
        '4hour': '(4??????)',
        '6hour': '(6??????)',
        '12hour': '(12??????)',
        '3day': '(3???)'
    },
    'en-us': {
        'line': '(Line)',
        '1min': '(1m)',
        '5min': '(5m)',
        '15min': '(15m)',
        '30min': '(30m)',
        '1hour': '(1h)',
        '1day': '(1d)',
        '1week': '(1w)',
        '3min': '(3m)',
        '2hour': '(2h)',
        '4hour': '(4h)',
        '6hour': '(6h)',
        '12hour': '(12h)',
        '3day': '(3d)'
    },
    'zh-tw': {
        'line': '(??????)',
        '1min': '(1??????)',
        '5min': '(5??????)',
        '15min': '(15??????)',
        '30min': '(30??????)',
        '1hour': '(1??????)',
        '1day': '(??????)',
        '1week': '(??????)',
        '3min': '(3??????)',
        '2hour': '(2??????)',
        '4hour': '(4??????)',
        '6hour': '(6??????)',
        '12hour': '(12??????)',
        '3day': '(3???)'
    }
};