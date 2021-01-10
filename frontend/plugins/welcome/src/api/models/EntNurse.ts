/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntNurseEdges,
    EntNurseEdgesFromJSON,
    EntNurseEdgesFromJSONTyped,
    EntNurseEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntNurse
 */
export interface EntNurse {
    /**
     * 
     * @type {EntNurseEdges}
     * @memberof EntNurse
     */
    edges?: EntNurseEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntNurse
     */
    id?: number;
    /**
     * NurseEmail holds the value of the "nurse_Email" field.
     * @type {string}
     * @memberof EntNurse
     */
    nurseEmail?: string;
    /**
     * NurseName holds the value of the "nurse_Name" field.
     * @type {string}
     * @memberof EntNurse
     */
    nurseName?: string;
    /**
     * NursePassword holds the value of the "nurse_Password" field.
     * @type {string}
     * @memberof EntNurse
     */
    nursePassword?: string;
    /**
     * NurseTel holds the value of the "nurse_Tel" field.
     * @type {string}
     * @memberof EntNurse
     */
    nurseTel?: string;
}

export function EntNurseFromJSON(json: any): EntNurse {
    return EntNurseFromJSONTyped(json, false);
}

export function EntNurseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntNurse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntNurseEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'nurseEmail': !exists(json, 'nurse_Email') ? undefined : json['nurse_Email'],
        'nurseName': !exists(json, 'nurse_Name') ? undefined : json['nurse_Name'],
        'nursePassword': !exists(json, 'nurse_Password') ? undefined : json['nurse_Password'],
        'nurseTel': !exists(json, 'nurse_Tel') ? undefined : json['nurse_Tel'],
    };
}

export function EntNurseToJSON(value?: EntNurse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntNurseEdgesToJSON(value.edges),
        'id': value.id,
        'nurse_Email': value.nurseEmail,
        'nurse_Name': value.nurseName,
        'nurse_Password': value.nursePassword,
        'nurse_Tel': value.nurseTel,
    };
}


